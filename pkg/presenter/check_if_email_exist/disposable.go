package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
)

type disposablePresenter struct {
	IsDisposable bool `json:"is_disposable"`
}

type disposablePreparer struct{}

func (disposablePreparer) CanPrepare(_ evmail.Address, result ev.ValidationResult, _ preparer.Options) bool {
	return result.ValidatorName() == ev.DisposableValidatorName
}

func (disposablePreparer) Prepare(_ evmail.Address, result ev.ValidationResult, _ preparer.Options) interface{} {
	return disposablePresenter{!result.IsValid()}
}
