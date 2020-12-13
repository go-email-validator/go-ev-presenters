package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-email-validator/pkg/ev/smtp_checker"
	"github.com/imdario/mergo"
	"strings"
)

type SMTPPresenter struct {
	CanConnectSmtp bool `json:"can_connect_smtp"`
	HasFullInbox   bool `json:"has_full_inbox"`
	IsCatchAll     bool `json:"is_catch_all"`
	IsDeliverable  bool `json:"is_deliverable"`
	IsDisabled     bool `json:"is_disabled"`
}

var SuccessSMTPPresenter = SMTPPresenter{true, false, true, true, false}
var FalseSMTPPresenter = SMTPPresenter{false, false, false, false, false}

type SMTPProcessor struct{}

func (s SMTPProcessor) CanProcess(_ email.EmailAddressInterface, result ev.ValidationResultInterface) bool {
	return result.ValidatorName() == ev.SMTPValidatorName
}

func (s SMTPProcessor) Process(_ email.EmailAddressInterface, result ev.ValidationResultInterface) interface{} {
	var presenter = SMTPPresenter{}
	var smtpError smtp_checker.SMTPError
	var errString string
	var ok bool

	for _, err := range result.Errors() {
		if smtpError, ok = err.(smtp_checker.SMTPError); !ok {
			continue
		}

		errString = smtpError.Err().Error()
		switch smtpError.Stage() {
		default:
			presenter = FalseSMTPPresenter
		case smtp_checker.RandomRCPTStage:
			presenter.IsCatchAll = false
		case smtp_checker.RCPTStage:
			presenter.IsDeliverable = false
			switch {
			case strings.Contains(errString, "disabled"),
				strings.Contains(errString, "discontinued"):
				presenter.IsDisabled = true
			case strings.Contains(errString, "full"),
				strings.Contains(errString, "insufficient"),
				strings.Contains(errString, "over quota"),
				strings.Contains(errString, "space"),
				strings.Contains(errString, "too many messages"):
				presenter.HasFullInbox = true
			case strings.Contains(errString, "the user you are trying to contact is receiving mail at a rate that"):
				presenter.IsDeliverable = true
			}
		}
	}

	mergo.Merge(&presenter, SuccessSMTPPresenter)
	return result
}
