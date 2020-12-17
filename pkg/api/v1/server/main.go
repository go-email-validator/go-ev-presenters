package server

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
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type EVApiV1 struct {
	presenter         presenter.Presenter
	preparersMatching map[v1.ResultType]preparer.Name
}

func (e EVApiV1) SingleValidation(_ context.Context, request *v1.EmailRequest) (*v1.EmailResponse, error) {
	var result v1.IsEmailResponse_Result

	present, err := e.presenter.SingleValidation(ev_email.EmailFromString(request.Email), e.preparersMatching[request.ResultType])
	if err != nil {
		return nil, err
	}

	switch v := present.(type) {
	case mailboxvalidator.DepPresenter:
		result = &v1.EmailResponse_MailBoxValidator{
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
		}
	default:
		ciee, ok := present.(check_if_email_exist.DepPresenter)
		if !ok {
			return nil, errors.New("invalid ResultType")
		}

		result = &v1.EmailResponse_CheckIfEmailExist{
			CheckIfEmailExist: &api_ciee.Result{
				Input:       ciee.Input,
				IsReachable: ciee.IsReachable.String(),
				Misc: &api_ciee.Misc{
					IsDisposable:  ciee.Misc.IsDisposable,
					IsRoleAccount: ciee.Misc.IsRoleAccount,
				},
				MX: &api_ciee.MX{
					AcceptsMail: ciee.MX.AcceptsMail,
					Records:     ciee.MX.Records,
				},
				SMTP: &api_ciee.SMTP{
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
		}
	}

	response := &v1.EmailResponse{Result: result}

	return response, err
}

func main() {
	server := grpc.NewServer()

	instance := &EVApiV1{
		presenter: presenter.NewPresenter(),
		preparersMatching: map[v1.ResultType]preparer.Name{
			v1.ResultType_CHECK_IF_EMAIL_EXIST: check_if_email_exist.Name,
			v1.ResultType_MAIL_BOX_VALIDATOR:   mailboxvalidator.Name,
		},
	}

	v1.RegisterEmailValidationServer(server, instance)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
