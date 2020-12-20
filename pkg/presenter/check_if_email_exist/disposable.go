package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
)

type disposablePresenter struct {
	IsDisposable bool `json:"is_disposable"`
}

type disposablePreparer struct{}

func (_ disposablePreparer) CanPrepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface, _ preparer.OptionsInterface) bool {
	return result.ValidatorName() == ev.DisposableValidatorName
}

func (_ disposablePreparer) Prepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface, _ preparer.OptionsInterface) interface{} {
	return disposablePresenter{!result.IsValid()}
}
