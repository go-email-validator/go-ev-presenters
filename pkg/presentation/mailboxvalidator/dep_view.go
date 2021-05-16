package mailboxvalidator

import (
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/converter"
	"time"
)

const (
	MBVTrue  = "True"
	MBVFalse = "False"
)

type DepPresentationForView struct {
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

func NewDepConverterForViewDefault() DepConverterForView {
	return NewDepConverterForView(NewDepConverterDefault())
}

func NewDepConverterForView(depConverter DepConverter) DepConverterForView {
	return DepConverterForView{depConverter}
}

type DepConverterForView struct {
	d DepConverter
}

func (d DepConverterForView) Can(email evmail.Address, result ev.ValidationResult, opts converter.Options) bool {
	return d.d.Can(email, result, opts)
}

// TODO add processing of "-" in mailbox validator, for example zxczxczxc@joycasinoru
func (d DepConverterForView) Convert(email evmail.Address, resultInterface ev.ValidationResult, opts converter.Options) interface{} {
	depPresentation := d.d.Convert(email, resultInterface, opts).(DepPresentation)

	return DepPresentationForView{
		EmailAddress:          depPresentation.EmailAddress,
		Domain:                depPresentation.Domain,
		IsFree:                FromBool(depPresentation.IsFree),
		IsSyntax:              FromBool(depPresentation.IsSyntax),
		IsDomain:              FromBool(depPresentation.IsDomain),
		IsSmtp:                FromBool(depPresentation.IsSmtp),
		IsVerified:            FromBool(depPresentation.IsVerified),
		IsServerDown:          FromBool(depPresentation.IsServerDown),
		IsGreylisted:          FromBool(depPresentation.IsGreylisted),
		IsDisposable:          FromBool(depPresentation.IsDisposable),
		IsSuppressed:          FromBool(depPresentation.IsSuppressed),
		IsRole:                FromBool(depPresentation.IsRole),
		IsHighRisk:            FromBool(depPresentation.IsHighRisk),
		IsCatchall:            FromBool(depPresentation.IsCatchall),
		MailboxvalidatorScore: fmt.Sprint(depPresentation.MailboxvalidatorScore),
		TimeTaken:             fmt.Sprint(depPresentation.TimeTaken.Round(time.Microsecond).Seconds()),
		Status:                FromBool(depPresentation.Status),
		CreditsAvailable:      depPresentation.CreditsAvailable,
		ErrorCode:             depPresentation.ErrorCode,
		ErrorMessage:          depPresentation.ErrorMessage,
	}
}
