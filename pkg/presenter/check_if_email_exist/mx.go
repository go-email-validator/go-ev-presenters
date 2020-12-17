package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
)

type mxPresenter struct {
	AcceptsMail bool     `json:"accepts_mail"`
	Records     []string `json:"records"`
}

type mxPreparer struct{}

func (s mxPreparer) CanPrepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface) bool {
	return result.ValidatorName() == ev.MXValidatorName
}

func (s mxPreparer) Prepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface) interface{} {
	mxResult := result.(ev.MXValidationResultInterface)
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
