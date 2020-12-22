package main

import (
	"context"
	"github.com/go-email-validator/go-ev-presenters/pkg/api/v1"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
	"strings"
	"sync"
)

const (
	domain      = "0.0.0.0"
	grpcPort    = ":50051"
	grpcAddress = domain + grpcPort
	httpPort    = ":50052"
	httpAddress = domain + httpPort
)

func main() {
	quit := make(chan bool)
	_, gr := runServer(quit)
	gr.Wait()
	defer closeServer(quit)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func runServer(quit chan bool) (*sync.WaitGroup, *errgroup.Group) {
	var grpcServer *grpc.Server
	var httpServer *http.Server
	go func() {
		if quit != nil {
			<-quit

			if httpServer != nil {
				httpServer.Close()
			}
			if grpcServer != nil {
				grpcServer.Stop()
			}
			quit <- true
		}
	}()
	var err error
	wg := new(sync.WaitGroup)
	wg.Add(2)

	instance := EVApiV1{
		presenter: presenter.NewMultiplePresentersDefault(),
		matching: map[v1.ResultType]preparer.Name{
			v1.ResultType_CHECK_IF_EMAIL_EXIST: check_if_email_exist.Name,
			v1.ResultType_MAIL_BOX_VALIDATOR:   mailboxvalidator.Name,
		},
	}

	grpcServer = grpc.NewServer()
	v1.RegisterEmailValidationServer(grpcServer, instance)

	var group = new(errgroup.Group)
	listener, err := net.Listen("tcp", grpcAddress)
	checkErr(err)
	group.Go(func() error {
		wg.Done()
		return grpcServer.Serve(listener)
	})

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{MarshalOptions: protojson.MarshalOptions{EmitUnpopulated: true, UseProtoNames: true}},
	))
	httpServer = &http.Server{Addr: httpAddress, Handler: mux}
	opts := []grpc.DialOption{
		grpc.WithInsecure(), grpc.WithBlock(),
	}

	ctx := context.Background()
	err = v1.RegisterEmailValidationHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	checkErr(err)
	group.Go(func() error {
		addr := httpServer.Addr
		if addr == "" {
			addr = ":http"
		}
		ln, err := net.Listen("tcp", addr)
		wg.Done()
		if err != nil {
			return err
		}
		err = httpServer.Serve(ln)
		if err != nil && !strings.Contains(err.Error(), "http: Server closed") {
			return err
		}
		return nil
	})

	go func() {
		err = group.Wait()
		checkErr(err)
	}()

	return wg, group
}

func closeServer(quit chan bool) {
	quit <- true
	<-quit
}
