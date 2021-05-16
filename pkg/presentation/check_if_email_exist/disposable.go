package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/converter"
)

type disposablePresentation struct {
	IsDisposable bool `json:"is_disposable"`
}

type disposablePreparer struct{}

func (disposablePreparer) CanPrepare(_ evmail.Address, result ev.ValidationResult, _ converter.Options) bool {
	return result.ValidatorName() == ev.DisposableValidatorName
}

func (disposablePreparer) Prepare(_ evmail.Address, result ev.ValidationResult, _ converter.Options) interface{} {
	return disposablePresentation{!result.IsValid()}
}
