package server

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp/smtp_client"
	"github.com/go-email-validator/go-ev-presenters/pkg/api/v1"
	"github.com/go-email-validator/go-ev-presenters/pkg/log"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"h12.io/socks"
	"io"
	"io/ioutil"
	"mime"
	"net"
	"net/http"
	"net/smtp"
	"strings"
	"sync"
)

func NewServer(opts Options) Server {
	return Server{
		opts: opts,
	}
}

type Server struct {
	grpcServer *grpc.Server
	httpServer *http.Server
	opts       Options
	waitGroup  sync.WaitGroup
}

func (s *Server) Start() error {
	if err := s.StartGRPC(); err != nil {
		return fmt.Errorf("starting gRPC server failed: %w", err)
	}

	if err := s.StartHTTP(); err != nil {
		return fmt.Errorf("starting HTTP server failed: %w", err)
	}

	return nil
}

func (s *Server) StartGRPC() error {
	l, err := net.Listen("tcp", s.opts.GRPC.Bind)

	if err != nil {
		return fmt.Errorf("create listener: %w", err)
	}

	var opts []grpc.ServerOption
	if s.opts.Auth.Key != "" {
		opts = append(opts, grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) {
			// From https://github.com/grpc-ecosystem/go-grpc-middleware/blob/master/auth/examples_test.go
			token := metautils.ExtractIncoming(ctx).Get("authorization")
			if token != s.opts.Auth.Key {
				return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: \"%v\"", token)
			}

			return context.Background(), nil
		})))
	}

	s.grpcServer = grpc.NewServer(opts...)

	var dialFunc evsmtp.DialFunc
	if s.opts.SMTPProxy != "" {
		dialFunc = func(addr string) (smtp_client.SMTPClient, error) {
			conn, err := socks.Dial(s.opts.SMTPProxy)("tcp", addr)
			if err != nil {
				log.Logger().Errorf("proxy create conn: %s", err)
				return nil, err
			}

			host, _, _ := net.SplitHostPort(addr)
			return smtp.NewClient(conn, host)
		}
	}

	v1.RegisterEmailValidationServer(s.grpcServer, defaultInstance(dialFunc))

	s.waitGroup.Add(1)
	go func() {
		log.Logger().Debugf("gRPC server is starting on %s", l.Addr())
		if err := s.grpcServer.Serve(l); err != nil {
			log.Logger().Errorf("gRPC server: %s", err)
		}
		s.waitGroup.Done()
		log.Logger().Debug("gRPC server stopped")
	}()

	return nil
}

func (s *Server) StartHTTP() error {
	if !s.opts.HTTP.Enable {
		return nil
	}

	l, err := net.Listen("tcp", s.opts.HTTP.Bind)

	if err != nil {
		return fmt.Errorf("create listener: %v", err)
	}

	mux := http.NewServeMux()

	err = s.addSwagger(mux)
	if err != nil {
		return err
	}

	gwmux := runtime.NewServeMux(s.opts.HTTP.MuxOptions...)
	mux.Handle("/", gwmux)

	ctx := context.Background()
	err = v1.RegisterEmailValidationHandlerFromEndpoint(ctx, gwmux, s.opts.GRPC.Bind, s.opts.HTTP.GRPCOptions)
	if err != nil {
		return err
	}

	s.httpServer = &http.Server{Addr: s.opts.HTTP.Bind, Handler: mux}

	s.waitGroup.Add(1)
	go func() {
		log.Logger().Debugf("HTTP server is starting on %s", l.Addr())
		err = s.httpServer.Serve(l)

		if err != nil && !strings.Contains(err.Error(), "http: Server closed") {
			log.Logger().Errorf("http server: %s", err)
		}
		s.waitGroup.Done()
		log.Logger().Debug("HTTP server stopped")
	}()

	return nil
}

func (s *Server) addSwagger(mux *http.ServeMux) error {
	swagger, err := ioutil.ReadFile(s.opts.HTTP.SwaggerPath)
	if err != nil {
		return err
	}
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, bytes.NewReader(swagger))
	})

	mime.AddExtensionType(".svg", "image/svg+xml")
	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(http.Dir("third_party/swagger-ui"))
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))

	return nil
}

func (s *Server) Wait() {
	s.waitGroup.Wait()
}

func (s *Server) Shutdown() {
	s.ShutdownHTTP()
	s.ShutdownGRPC()
}

func (s *Server) ShutdownGRPC() {
	ctx, cancel := context.WithTimeout(context.Background(), s.opts.GRPC.ShutdownTimeout)
	defer cancel()
	ch := make(chan bool, 1)
	go func() {
		s.grpcServer.GracefulStop()
		ch <- true
	}()
	for {
		select {
		case <-ctx.Done():
			s.grpcServer.Stop()
		case <-ch:
			s.grpcServer.Stop()
		}
	}
}

func (s *Server) ShutdownHTTP() {
	if !s.opts.HTTP.Enable {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.opts.HTTP.ShutdownTimeout)
	defer cancel()
	s.httpServer.Shutdown(ctx)
}
