package server

import (
	"context"
	"errors"
	v1 "github.com/go-email-validator/go-ev-presenters/pkg/api/v1"
	apiciee "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/check_if_email_exist"
	apimbv "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/mailboxvalidator"
	api_prompt_email_verification "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/prompt_email_verification_api"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/prompt_email_verification_api"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type EVApiV1 struct {
	presenter presenter.MultiplePresenter
	matching  map[v1.EmailRequest_ResultType]preparer.Name
	v1.UnsafeEmailValidationServer
}

func (e EVApiV1) SingleValidation(_ context.Context, request *v1.EmailRequest) (*v1.EmailResponse, error) {
	var response *v1.EmailResponse

	present, err := e.presenter.SingleValidation(request.Email, e.matching[request.ResultType])
	if err != nil {
		return nil, err
	}

	switch v := present.(type) {
	case mailboxvalidator.DepPresenterForView:
		response = &v1.EmailResponse{Result: &v1.EmailResponse_MailBoxValidator{
			MailBoxValidator: &apimbv.Result{
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
	case prompt_email_verification_api.DepPresenter:
		response = &v1.EmailResponse{Result: &v1.EmailResponse_PromptEmailVerificationApi{
			PromptEmailVerificationApi: &api_prompt_email_verification.Result{
				CanConnectSmtp: v.CanConnectSmtp,
				Email:          v.Email,
				IsCatchAll:     v.IsCatchAll,
				IsDeliverable:  v.IsDeliverable,
				IsDisabled:     v.IsDisabled,
				IsDisposable:   v.IsDisposable,
				IsInboxFull:    v.IsInboxFull,
				IsRoleAccount:  v.IsRoleAccount,
				MxRecords: &api_prompt_email_verification.Result_MX{
					AcceptsMail: v.MxRecords.AcceptsMail,
					Records:     v.MxRecords.Records,
				},
				SyntaxValid: v.SyntaxValid,
				Message:     v.Message,
			}},
		}
	case check_if_email_exist.DepPresenter:
		var address *wrappers.StringValue
		if v.Syntax.Address == nil {
			address = nil
		} else {
			address = &wrappers.StringValue{Value: *v.Syntax.Address}
		}
		response = &v1.EmailResponse{Result: &v1.EmailResponse_CheckIfEmailExist{
			CheckIfEmailExist: &apiciee.Result{
				Input:       v.Input,
				IsReachable: v.IsReachable.String(),
				Misc: &apiciee.Misc{
					IsDisposable:  v.Misc.IsDisposable,
					IsRoleAccount: v.Misc.IsRoleAccount,
				},
				Mx: &apiciee.MX{
					AcceptsMail: v.MX.AcceptsMail,
					Records:     v.MX.Records,
				},
				Smtp: &apiciee.SMTP{
					CanConnectSmtp: v.SMTP.CanConnectSmtp,
					HasFullInbox:   v.SMTP.HasFullInbox,
					IsCatchAll:     v.SMTP.IsCatchAll,
					IsDeliverable:  v.SMTP.IsDeliverable,
					IsDisabled:     v.SMTP.IsDisabled,
				},
				Syntax: &apiciee.Syntax{
					Address:       address,
					Username:      v.Syntax.Username,
					Domain:        v.Syntax.Domain,
					IsValidSyntax: v.Syntax.IsValidSyntax,
				},
				Error: v.Error,
			},
		},
		}
	default:
		return nil, errors.New("invalid ResultType")
	}

	return response, err
}
