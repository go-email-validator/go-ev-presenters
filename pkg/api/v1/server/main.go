package main

import (
	"context"
	"errors"
	"github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/api/v1"
	api_ciee "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/check_if_email_exist"
	api_mbv "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/mailboxvalidator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
)

const (
	domain      = "0.0.0.0"
	grpcPort    = ":50051"
	grpcAddress = domain + grpcPort
	httpPort    = ":50052"
	httpAddress = domain + httpPort
)

type EVApiV1 struct {
	presenter         presenter.Presenter
	preparersMatching map[v1.ResultType]preparer.Name
	v1.UnsafeEmailValidationServer
}

func (e EVApiV1) SingleValidation(_ context.Context, request *v1.EmailRequest) (*v1.EmailResponse, error) {
	var response v1.EmailResponse

	present, err := e.presenter.SingleValidation(ev_email.EmailFromString(request.Email), e.preparersMatching[request.ResultType])
	if err != nil {
		return nil, err
	}

	switch v := present.(type) {
	case mailboxvalidator.DepPresenter:
		response = v1.EmailResponse{Result: &v1.EmailResponse_MailBoxValidator{
			MailBoxValidator: &api_mbv.Result{
				EmailAddress:          v.EmailAddress,
				Domain:                v.Domain,
				IsFree:                v.IsFree,
				IsSyntax:              v.IsSyntax,
				IsDomain:              v.IsDomain,
				IsSmtp:                v.IsSmtp,
				IsVerified:            v.IsVerified,
				IsServerDown:          v.IsServerDown,
				IsGreylisted:          v.IsGreylisted,
				IsDisposable:          v.IsDisposable,
				IsSuppressed:          v.IsSuppressed,
				IsRole:                v.IsRole,
				IsHighRisk:            v.IsHighRisk,
				IsCatchall:            v.IsCatchall,
				MailboxvalidatorScore: v.MailboxvalidatorScore,
				TimeTaken:             v.TimeTaken,
				Status:                v.Status,
				CreditsAvailable:      v.CreditsAvailable,
				ErrorCode:             v.ErrorCode,
				ErrorMessage:          v.ErrorMessage,
			},
		},
		}
	default:
		ciee, ok := present.(check_if_email_exist.DepPresenter)
		if !ok {
			return nil, errors.New("invalid ResultType")
		}

		response = v1.EmailResponse{Result: &v1.EmailResponse_CheckIfEmailExist{
			CheckIfEmailExist: &api_ciee.Result{
				Input:       ciee.Input,
				IsReachable: ciee.IsReachable.String(),
				Misc: &api_ciee.Misc{
					IsDisposable:  ciee.Misc.IsDisposable,
					IsRoleAccount: ciee.Misc.IsRoleAccount,
				},
				Mx: &api_ciee.MX{
					AcceptsMail: ciee.MX.AcceptsMail,
					Records:     ciee.MX.Records,
				},
				Smtp: &api_ciee.SMTP{
					CanConnectSmtp: ciee.SMTP.CanConnectSmtp,
					HasFullInbox:   ciee.SMTP.HasFullInbox,
					IsCatchAll:     ciee.SMTP.IsCatchAll,
					IsDeliverable:  ciee.SMTP.IsDeliverable,
					IsDisabled:     ciee.SMTP.IsDisabled,
				},
				Syntax: &api_ciee.Syntax{
					Username:      ciee.Syntax.Username,
					Domain:        ciee.Syntax.Domain,
					IsValidSyntax: ciee.Syntax.IsValidSyntax,
				},
			},
		},
		}
	}

	return &response, err
}

func main() {
	var err error

	instance := EVApiV1{
		presenter: presenter.NewPresenter(),
		preparersMatching: map[v1.ResultType]preparer.Name{
			v1.ResultType_CHECK_IF_EMAIL_EXIST: check_if_email_exist.Name,
			v1.ResultType_MAIL_BOX_VALIDATOR:   mailboxvalidator.Name,
		},
	}

	grpcServer := grpc.NewServer()
	v1.RegisterEmailValidationServer(grpcServer, instance)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var group errgroup.Group
	group.Go(func() error {
		listener, err := net.Listen("tcp", grpcAddress)
		if err != nil {
			log.Fatal("Unable to create grpc listener:", err)
		}
		return grpcServer.Serve(listener)
	})

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{MarshalOptions: protojson.MarshalOptions{EmitUnpopulated: true}},
	))
	opts := []grpc.DialOption{
		grpc.WithInsecure(), grpc.WithBlock(),
	}
	group.Go(func() error {
		return v1.RegisterEmailValidationHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	})

	group.Go(func() error {
		return http.ListenAndServe(httpAddress, mux)
	})

	err = group.Wait()
	if err != nil {
		panic(err)
	}

	return
}