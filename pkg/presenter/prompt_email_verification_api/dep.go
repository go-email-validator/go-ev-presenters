package prompt_email_verification_api

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
)

//go:generate go run cmd/dep_test_generator/gen.go

const (
	Name               preparer.Name = "PromptEmailVerificationApi"
	ErrorSyntaxInvalid               = "Invalid email syntax."
)

type mx struct {
	AcceptsMail bool     `json:"accepts_mail"`
	Records     []string `json:"records"`
}

// https://promptapi.com/marketplace/description/email_verification-api
type DepPresenter struct {
	Email          string `json:"email"`
	SyntaxValid    bool   `json:"syntax_valid"`
	IsDisposable   bool   `json:"is_disposable"`
	IsRoleAccount  bool   `json:"is_role_account"`
	IsCatchAll     bool   `json:"is_catch_all"`
	IsDeliverable  bool   `json:"is_deliverable"`
	CanConnectSmtp bool   `json:"can_connect_smtp"`
	IsInboxFull    bool   `json:"is_inbox_full"`
	IsDisabled     bool   `json:"is_disabled"`
	MxRecords      mx     `json:"mx_records"`
	Message        string `json:"message"`
}

func NewDepPreparerDefault() DepPreparer {
	return NewDepPreparer()
}

func NewDepPreparer() DepPreparer {
	return DepPreparer{}
}

type DepPreparer struct{}

func (DepPreparer) CanPrepare(_ evmail.Address, result ev.ValidationResult, _ preparer.Options) bool {
	return result.ValidatorName() == ev.DepValidatorName
}

func (d DepPreparer) Prepare(email evmail.Address, resultInterface ev.ValidationResult, _ preparer.Options) (result interface{}) {
	var message string
	depResult := resultInterface.(ev.DepValidationResult)
	validationResults := depResult.GetResults()
	mxResult := validationResults[ev.MXValidatorName].(ev.MXValidationResult)

	smtpPresenter := common.SMTPPreparer{}.Prepare(email, validationResults[ev.SMTPValidatorName], nil).(common.SmtpPresenter)

	Email := email.String()
	isSyntaxValid := validationResults[ev.SyntaxValidatorName].IsValid()
	if !isSyntaxValid && len(Email) == 0 {
		message = ErrorSyntaxInvalid
	}

	depPresenter := DepPresenter{
		Email:          Email,
		SyntaxValid:    isSyntaxValid,
		IsDisposable:   !validationResults[ev.DisposableValidatorName].IsValid(),
		IsRoleAccount:  !validationResults[ev.RoleValidatorName].IsValid(),
		IsCatchAll:     smtpPresenter.IsCatchAll,
		IsDeliverable:  smtpPresenter.IsDeliverable,
		CanConnectSmtp: smtpPresenter.CanConnectSmtp,
		IsInboxFull:    smtpPresenter.HasFullInbox,
		IsDisabled:     smtpPresenter.IsDisabled,
		MxRecords: mx{
			AcceptsMail: mxResult.IsValid(),
			Records:     common.MX2String(mxResult.MX()),
		},
		Message: message,
	}

	return depPresenter
}

func NewDepValidator() ev.Validator {
	return ev.NewDepBuilder(nil).Set(
		ev.SyntaxValidatorName,
		common.NewSyntaxValidator(),
	).Build()
}
