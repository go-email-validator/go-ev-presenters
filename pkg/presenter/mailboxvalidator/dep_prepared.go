package mailboxvalidator

import (
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"time"
)

const (
	MBVTrue  = "True"
	MBVFalse = "False"
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

func FromBool(value bool) string {
	if value {
		return MBVTrue
	}
	return MBVFalse
}

func ToBool(value string) bool {
	return value == MBVTrue
}

func NewDepPreparerForViewDefault() DepPreparerForView {
	return NewDepPreparerForView(NewDepPreparerDefault())
}

func NewDepPreparerForView(preparer DepPreparer) DepPreparerForView {
	return DepPreparerForView{preparer}
}

type DepPreparerForView struct {
	d DepPreparer
}

func (d DepPreparerForView) CanPrepare(email evmail.Address, result ev.ValidationResult, opts preparer.Options) bool {
	return d.d.CanPrepare(email, result, opts)
}

// TODO add processing of "-" in mailbox validator, for example zxczxczxc@joycasinoru
func (d DepPreparerForView) Prepare(email evmail.Address, resultInterface ev.ValidationResult, opts preparer.Options) interface{} {
	depPresenter := d.d.Prepare(email, resultInterface, opts).(DepPresenter)

	return DepPresenterForView{
		EmailAddress:          depPresenter.EmailAddress,
		Domain:                depPresenter.Domain,
		IsFree:                FromBool(depPresenter.IsFree),
		IsSyntax:              FromBool(depPresenter.IsSyntax),
		IsDomain:              FromBool(depPresenter.IsDomain),
		IsSmtp:                FromBool(depPresenter.IsSmtp),
		IsVerified:            FromBool(depPresenter.IsVerified),
		IsServerDown:          FromBool(depPresenter.IsServerDown),
		IsGreylisted:          FromBool(depPresenter.IsGreylisted),
		IsDisposable:          FromBool(depPresenter.IsDisposable),
		IsSuppressed:          FromBool(depPresenter.IsSuppressed),
		IsRole:                FromBool(depPresenter.IsRole),
		IsHighRisk:            FromBool(depPresenter.IsHighRisk),
		IsCatchall:            FromBool(depPresenter.IsCatchall),
		MailboxvalidatorScore: fmt.Sprint(depPresenter.MailboxvalidatorScore),
		TimeTaken:             fmt.Sprint(depPresenter.TimeTaken.Round(time.Microsecond).Seconds()),
		Status:                FromBool(depPresenter.Status),
		CreditsAvailable:      depPresenter.CreditsAvailable,
		ErrorCode:             depPresenter.ErrorCode,
		ErrorMessage:          depPresenter.ErrorMessage,
	}
}
