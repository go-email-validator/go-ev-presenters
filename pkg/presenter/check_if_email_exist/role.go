package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
)

type rolePresenter struct {
	IsRoleAccount bool `json:"is_role_account"`
}

type rolePreparer struct{}

func (s rolePreparer) CanPrepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface) bool {
	return result.ValidatorName() == ev.RoleValidatorName
}

func (s rolePreparer) Prepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface) interface{} {
	return rolePresenter{!result.IsValid()}
}
