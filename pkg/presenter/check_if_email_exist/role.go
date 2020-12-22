package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
)

type rolePresenter struct {
	IsRoleAccount bool `json:"is_role_account"`
}

type rolePreparer struct{}

func (_ rolePreparer) CanPrepare(_ email.EmailAddress, result ev.ValidationResult, _ preparer.Options) bool {
	return result.ValidatorName() == ev.RoleValidatorName
}

func (_ rolePreparer) Prepare(_ email.EmailAddress, result ev.ValidationResult, _ preparer.Options) interface{} {
	return rolePresenter{!result.IsValid()}
}
