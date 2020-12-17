package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
)

type syntaxPresenter struct {
	Username      string `json:"username"`
	Domain        string `json:"domain"`
	IsValidSyntax bool   `json:"is_valid_syntax"`
}

type SyntaxPreparer struct{}

func (s SyntaxPreparer) CanPrepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface) bool {
	return result.ValidatorName() == ev.SyntaxValidatorName
}

func (s SyntaxPreparer) Prepare(email email.EmailAddressInterface, result ev.ValidationResultInterface) interface{} {
	return syntaxPresenter{
		email.Username(),
		email.Domain(),
		result.IsValid(),
	}
}
