package server

import (
	"context"
	"fmt"
	"github.com/go-email-validator/go-ev-presenters/pkg/api/v1"
	"github.com/go-email-validator/go-ev-presenters/pkg/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
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

	s.grpcServer = grpc.NewServer()
	v1.RegisterEmailValidationServer(s.grpcServer, s.opts.GRPC.Server)

	s.waitGroup.Add(1)
	go func() {
		log.Logger().Debug("gRPC server is starting")
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

	mux := runtime.NewServeMux(s.opts.HTTP.MuxOptions...)
	s.httpServer = &http.Server{Addr: s.opts.HTTP.Bind, Handler: mux}

	ctx := context.Background()
	err = v1.RegisterEmailValidationHandlerFromEndpoint(ctx, mux, s.opts.GRPC.Bind, s.opts.HTTP.GRPCOptions)
	if err != nil {
		return err
	}

	s.waitGroup.Add(1)
	go func() {
		log.Logger().Debug("HTTP server is starting")
		err = s.httpServer.Serve(l)

		if err != nil && !strings.Contains(err.Error(), "http: Server closed") {
			log.Logger().Errorf("http server: %s", err)
		}
		s.waitGroup.Done()
		log.Logger().Debug("HTTP server stopped")
	}()

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
