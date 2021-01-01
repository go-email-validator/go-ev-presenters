package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
)

type mxPresenter struct {
	AcceptsMail bool     `json:"accepts_mail"`
	Records     []string `json:"records"`
}

type mxPreparer struct{}

func (mxPreparer) CanPrepare(_ evmail.Address, result ev.ValidationResult, _ preparer.Options) bool {
	return result.ValidatorName() == ev.MXValidatorName
}

func (mxPreparer) Prepare(_ evmail.Address, result ev.ValidationResult, _ preparer.Options) interface{} {
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
