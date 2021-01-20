package common

import (
	"errors"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp"
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

var (
	WithoutErrsSMTPPresenter = SmtpPresenter{
		CanConnectSmtp: true,
		HasFullInbox:   false,
		IsCatchAll:     true,
		IsDeliverable:  true,
		IsDisabled:     false,
		IsGreyListed:   false,
	}
	FalseSMTPPresenter = SmtpPresenter{
		CanConnectSmtp: false,
		HasFullInbox:   false,
		IsCatchAll:     false,
		IsDeliverable:  false,
		IsDisabled:     false,
		IsGreyListed:   false,
	}
)

type SMTPPreparer struct{}

func (SMTPPreparer) CanPrepare(_ evmail.Address, result ev.ValidationResult, _ preparer.Options) bool {
	return result.ValidatorName() == ev.SMTPValidatorName
}

func (SMTPPreparer) Prepare(_ evmail.Address, result ev.ValidationResult, _ preparer.Options) interface{} {
	var presenter = WithoutErrsSMTPPresenter
	var errString string
	var errCode int
	var smtpError evsmtp.Error
	var depError *ev.DepsError

	errs := result.Errors()
	errs = append(errs, result.Warnings()...)
	for _, err := range errs {
		if !errors.As(err, &smtpError) {
			if errors.As(err, &depError) {
				return FalseSMTPPresenter
			}
			continue
		}

		sourceErr := errors.Unwrap(smtpError)
		errString = strings.ToLower(sourceErr.Error())

		errCode = 0
		switch v := sourceErr.(type) {
		case *textproto.Error:
			errCode = v.Code
		}
		if strings.Contains(errString, "greylist") {
			presenter.IsGreyListed = true
		}

		switch smtpError.Stage() {
		case evsmtp.ConnectionStage:
			presenter = FalseSMTPPresenter
		case evsmtp.HelloStage,
			evsmtp.AuthStage,
			evsmtp.MailStage:
			presenter.IsDeliverable = false
		case evsmtp.RandomRCPTStage:
			presenter.IsCatchAll = false
		case evsmtp.RCPTsStage:
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
