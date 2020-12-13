package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
)

type DisposablePresenter struct {
	IsDisposable bool `json:"is_disposable"`
}

type DisposableProcessor struct{}

func (s DisposableProcessor) CanProcess(_ email.EmailAddressInterface, result ev.ValidationResultInterface) bool {
	return result.ValidatorName() == ev.DisposableValidatorName
}

func (s DisposableProcessor) Process(_ email.EmailAddressInterface, result ev.ValidationResultInterface) interface{} {
	return DisposablePresenter{result.IsValid()}
}
