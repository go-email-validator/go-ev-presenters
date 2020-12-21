package mailboxvalidator

import (
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"time"
)

type DepPresenterForView struct {
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

func NewDepPreparerForViewDefault() DepPreparerForView {
	return NewDepPreparerForView(NewDepPreparerDefault())
}

func NewDepPreparerForView(preparer DepPreparer) DepPreparerForView {
	return DepPreparerForView{preparer}
}

type DepPreparerForView struct {
	DepPreparer
}

func (_ DepPreparerForView) CanPrepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface, opts preparer.OptionsInterface) bool {
	_, ok := opts.(preparer.TimeOptions)
	return ok && result.ValidatorName() == ev.DepValidatorName
}

func (d DepPreparerForView) Prepare(email email.EmailAddressInterface, resultInterface ev.ValidationResultInterface, opts preparer.OptionsInterface) interface{} {
	depPresenter := d.Prepare(email, resultInterface, opts).(DepPresenter)

	return DepPresenterForView{
		EmailAddress:          depPresenter.EmailAddress,
		Domain:                depPresenter.Domain,
		IsFree:                boolToString(depPresenter.IsFree),
		IsSyntax:              boolToString(depPresenter.IsSyntax),
		IsDomain:              boolToString(depPresenter.IsDomain),
		IsSmtp:                boolToString(depPresenter.IsSmtp),
		IsVerified:            boolToString(depPresenter.IsVerified),
		IsServerDown:          boolToString(depPresenter.IsServerDown),
		IsGreylisted:          boolToString(depPresenter.IsGreylisted),
		IsDisposable:          boolToString(depPresenter.IsDisposable),
		IsSuppressed:          boolToString(depPresenter.IsSuppressed),
		IsRole:                boolToString(depPresenter.IsRole),
		IsHighRisk:            boolToString(depPresenter.IsHighRisk),
		IsCatchall:            boolToString(depPresenter.IsCatchall),
		MailboxvalidatorScore: fmt.Sprint(depPresenter.MailboxvalidatorScore),
		TimeTaken:             fmt.Sprint(depPresenter.TimeTaken.Round(time.Microsecond).Seconds()),
		Status:                boolToString(depPresenter.Status),
		CreditsAvailable:      depPresenter.CreditsAvailable,
		ErrorCode:             depPresenter.ErrorCode,
		ErrorMessage:          depPresenter.ErrorMessage,
	}
}
