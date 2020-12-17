package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
)

type disposablePresenter struct {
	IsDisposable bool `json:"is_disposable"`
}

type disposablePreparer struct{}

func (s disposablePreparer) CanPrepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface) bool {
	return result.ValidatorName() == ev.DisposableValidatorName
}

func (s disposablePreparer) Prepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface) interface{} {
	return disposablePresenter{!result.IsValid()}
}
