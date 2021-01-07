package server

import (
	v1 "github.com/go-email-validator/go-ev-presenters/pkg/api/v1"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/prompt_email_verification_api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"time"
)

const (
	GRPCDefaultHost = "0.0.0.0:50051"
	HTTPDefaultHost = "0.0.0.0:50052"
)

func defaultInstance() v1.EmailValidationServer {
	return EVApiV1{
		presenter: getPresenter(),
		matching: map[v1.ResultType]preparer.Name{
			v1.ResultType_CHECK_IF_EMAIL_EXIST:          check_if_email_exist.Name,
			v1.ResultType_MAIL_BOX_VALIDATOR:            mailboxvalidator.Name,
			v1.ResultType_PROMPT_EMAIL_VERIFICATION_API: prompt_email_verification_api.Name,
		},
	}
}

var getPresenter = presenter.NewMultiplePresentersDefault

func NewOptions() Options {
	return Options{
		GRPC: NewGRPCOptions(),
		HTTP: NewHTTPOptions(),
	}
}

type Options struct {
	GRPC GRPCOptions
	HTTP HTTPOptions
}

var shutDownTimeout = 1 * time.Second

func NewGRPCOptions() GRPCOptions {
	return GRPCOptions{
		Bind:            GRPCDefaultHost,
		Server:          defaultInstance(),
		ShutdownTimeout: shutDownTimeout,
	}
}

type GRPCOptions struct {
	Bind            string
	Server          v1.EmailValidationServer
	ShutdownTimeout time.Duration
}

func NewHTTPOptions() HTTPOptions {
	return HTTPOptions{
		Bind: HTTPDefaultHost,
		MuxOptions: []runtime.ServeMuxOption{runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{MarshalOptions: protojson.MarshalOptions{EmitUnpopulated: true, UseProtoNames: true}},
		)},
		GRPCOptions: []grpc.DialOption{
			grpc.WithInsecure(), grpc.WithBlock(),
		},
		ShutdownTimeout: shutDownTimeout,
		Enable:          true,
	}
}

type HTTPOptions struct {
	Enable          bool
	Bind            string
	MuxOptions      []runtime.ServeMuxOption
	GRPCOptions     []grpc.DialOption
	ShutdownTimeout time.Duration
}
