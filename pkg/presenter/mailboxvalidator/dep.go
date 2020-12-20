package mailboxvalidator

import (
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"time"
)

const (
	Name = "MailBoxValidator"

	MissingParameter    = 100
	ApiKeyNotFound      = MissingParameter + iota
	ApiKeyDisabled      = MissingParameter + iota
	ApiKeyExpired       = MissingParameter + iota
	InsufficientCredits = MissingParameter + iota
	UnknownError        = MissingParameter + iota
)

type DepPresenter struct {
	EmailAddress          string `json:"email_address"`
	Domain                string `json:"domain"`
	IsFree                string `json:"is_free"`
	IsSyntax              string `json:"is_syntax"`
	IsDomain              string `json:"is_domain"`
	IsSmtp                string `json:"is_smtp"`
	IsVerified            string `json:"is_verified"`
	IsServerDown          string `json:"is_server_down"`
	IsGreylisted          string `json:"is_greylisted"`
	IsDisposable          string `json:"is_disposable"`
	IsSuppressed          string `json:"is_suppressed"`
	IsRole                string `json:"is_role"`
	IsHighRisk            string `json:"is_high_risk"`
	IsCatchall            string `json:"is_catchall"`
	MailboxvalidatorScore string `json:"mailboxvalidator_score"`
	TimeTaken             string `json:"time_taken"`
	Status                string `json:"status"`
	CreditsAvailable      uint32 `json:"credits_available"`
	ErrorCode             string `json:"error_code"`
	ErrorMessage          string `json:"error_message"`
}

func boolToString(value bool) string {
	if value {
		return "True"
	}
	return "False"
}

func NewDepPreparer() DepPreparer {
	return DepPreparer{}
}

type DepPreparer struct{}

func (_ DepPreparer) CanPrepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface, opts preparer.OptionsInterface) bool {
	_, ok := opts.(preparer.TimeOptions)
	return ok && result.ValidatorName() == ev.DepValidatorName
}

func (_ DepPreparer) Prepare(email email.EmailAddressInterface, resultInterface ev.ValidationResultInterface, opts preparer.OptionsInterface) interface{} {
	optsTime := opts.(preparer.TimeOptions)
	result := resultInterface.(ev.DepValidatorResult)

	depPresenter := DepPresenter{
		EmailAddress:          email.String(),
		Domain:                email.Domain(),
		IsFree:                boolToString(result.GetResults()[ev.FreeValidatorName].IsValid()),
		IsSyntax:              boolToString(result.GetResults()[ev.SyntaxValidatorName].IsValid()),
		IsDomain:              boolToString(result.GetResults()[ev.MXValidatorName].IsValid()),
		IsSmtp:                "",
		IsVerified:            "",
		IsServerDown:          "",
		IsGreylisted:          "",
		IsDisposable:          boolToString(result.GetResults()[ev.DisposableValidatorName].IsValid()),
		IsSuppressed:          "",
		IsRole:                boolToString(result.GetResults()[ev.RoleValidatorName].IsValid()),
		IsHighRisk:            "",
		IsCatchall:            "",
		MailboxvalidatorScore: "",
		TimeTaken:             fmt.Sprint(optsTime.ExecutedTime().Round(time.Microsecond).Seconds()),
		Status:                "",
		CreditsAvailable:      0,
		ErrorCode:             "",
		ErrorMessage:          "",
	}

	return depPresenter
}
