package common

import (
	"errors"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-email-validator/pkg/ev/smtp_checker"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"net/textproto"
	"strings"
)

type SmtpPresenter struct {
	CanConnectSmtp bool `json:"can_connect_smtp"`
	HasFullInbox   bool `json:"has_full_inbox"`
	IsCatchAll     bool `json:"is_catch_all"`
	IsDeliverable  bool `json:"is_deliverable"`
	IsDisabled     bool `json:"is_disabled"`
	IsGreyListed   bool `json:"is_grey_listed"`
}

var WithoutErrsSMTPPresenter = SmtpPresenter{
	CanConnectSmtp: true,
	HasFullInbox:   false,
	IsCatchAll:     true,
	IsDeliverable:  true,
	IsDisabled:     false,
	IsGreyListed:   false,
}
var FalseSMTPPresenter = SmtpPresenter{
	CanConnectSmtp: false,
	HasFullInbox:   false,
	IsCatchAll:     false,
	IsDeliverable:  false,
	IsDisabled:     false,
	IsGreyListed:   false,
}

type SMTPPreparer struct{}

func (_ SMTPPreparer) CanPrepare(_ email.EmailAddress, result ev.ValidationResult, _ preparer.Options) bool {
	return result.ValidatorName() == ev.SMTPValidatorName
}

func (_ SMTPPreparer) Prepare(_ email.EmailAddress, result ev.ValidationResult, _ preparer.Options) interface{} {
	var presenter = WithoutErrsSMTPPresenter
	var smtpError smtp_checker.SMTPError
	var errString string
	var errCode int

	errs := result.Errors()
	errs = append(errs, result.Warnings()...)
	for _, err := range errs {
		if !errors.As(err, &smtpError) {
			continue
		}

		errString = strings.ToLower(smtpError.Err().Error())
		errCode = smtpError.Err().(*textproto.Error).Code
		if strings.Contains(errString, "greylist") {
			presenter.IsGreyListed = true
		}

		switch smtpError.Stage() {
		case smtp_checker.ConnectionStage:
			presenter = FalseSMTPPresenter
		case smtp_checker.HelloStage,
			smtp_checker.AuthStage,
			smtp_checker.MailStage:
			presenter.IsDeliverable = false
		case smtp_checker.RandomRCPTStage:
			presenter.IsCatchAll = false
		case smtp_checker.RCPTStage:
			presenter.IsDeliverable = false
			switch {
			case strings.Contains(errString, "disabled") ||
				strings.Contains(errString, "discontinued"):
				presenter.IsDisabled = true
			case errCode == 452 && (strings.Contains(errString, "full") ||
				strings.Contains(errString, "insufficient") ||
				strings.Contains(errString, "over quota") ||
				strings.Contains(errString, "space") ||
				strings.Contains(errString, "too many messages")):

				presenter.HasFullInbox = true
			case strings.Contains(errString, "the user you are trying to contact is receiving mail at a rate that"):
				presenter.IsDeliverable = true
			}
		}
	}

	return presenter
}
