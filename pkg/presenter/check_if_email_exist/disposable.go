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

func (_ disposablePreparer) CanPrepare(_ email.EmailAddress, result ev.ValidationResult, _ preparer.Options) bool {
	return result.ValidatorName() == ev.DisposableValidatorName
}

func (_ disposablePreparer) Prepare(_ email.EmailAddress, result ev.ValidationResult, _ preparer.Options) interface{} {
	return disposablePresenter{!result.IsValid()}
}
