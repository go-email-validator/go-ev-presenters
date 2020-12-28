// TODO move to go-email-validator, do emailRegex changeable
package common

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-email-validator/pkg/ev/utils"
	"regexp"
)

var emailRegex = regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)])")

func NewSyntaxValidator() ev.Validator {
	return syntaxValidator{}
}

type syntaxValidator struct {
	ev.AValidatorWithoutDeps
}

func (_ syntaxValidator) Validate(email ev_email.EmailAddress, _ ...ev.ValidationResult) ev.ValidationResult {
	if emailRegex.MatchString(email.String()) {
		return ev.NewValidValidatorResult(ev.SyntaxValidatorName)
	}
	return syntaxGetError()
}

func syntaxGetError() ev.ValidationResult {
	return ev.NewValidatorResult(false, utils.Errs(ev.SyntaxError{}), nil, ev.SyntaxValidatorName)
}
