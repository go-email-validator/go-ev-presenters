package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
)

type syntaxPresenter struct {
	Username      string `json:"username"`
	Domain        string `json:"domain"`
	IsValidSyntax bool   `json:"is_valid_syntax"`
}

type SyntaxPreparer struct{}

func (_ SyntaxPreparer) CanPrepare(_ email.EmailAddress, result ev.ValidationResult, _ preparer.Options) bool {
	return result.ValidatorName() == ev.SyntaxValidatorName
}

func (_ SyntaxPreparer) Prepare(email email.EmailAddress, result ev.ValidationResult, _ preparer.Options) interface{} {
	return syntaxPresenter{
		email.Username(),
		email.Domain(),
		result.IsValid(),
	}
}
