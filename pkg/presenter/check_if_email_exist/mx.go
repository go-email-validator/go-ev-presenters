package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
)

type mxPresenter struct {
	AcceptsMail bool     `json:"accepts_mail"`
	Records     []string `json:"records"`
}

type mxPreparer struct{}

func (_ mxPreparer) CanPrepare(_ email.EmailAddress, result ev.ValidationResult, _ preparer.Options) bool {
	return result.ValidatorName() == ev.MXValidatorName
}

func (_ mxPreparer) Prepare(_ email.EmailAddress, result ev.ValidationResult, _ preparer.Options) interface{} {
	mxResult := result.(ev.MXValidationResult)
	lenMX := len(mxResult.MX())
	records := make([]string, lenMX)

	for i, mx := range mxResult.MX() {
		records[i] = mx.Host
	}

	return mxPresenter{
		lenMX > 0,
		records,
	}
}
