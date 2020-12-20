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

func (_ rolePreparer) CanPrepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface, _ preparer.OptionsInterface) bool {
	return result.ValidatorName() == ev.RoleValidatorName
}

func (_ rolePreparer) Prepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface, _ preparer.OptionsInterface) interface{} {
	return rolePresenter{!result.IsValid()}
}
