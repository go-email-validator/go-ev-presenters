package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
)

type RolePresenter struct {
	IsRoleAccount bool `json:"is_role_account"`
}

type RoleProcessor struct{}

func (s RoleProcessor) CanProcess(_ email.EmailAddressInterface, result ev.ValidationResultInterface) bool {
	return result.ValidatorName() == ev.RoleValidatorName
}

func (s RoleProcessor) Process(_ email.EmailAddressInterface, result ev.ValidationResultInterface) interface{} {
	return RolePresenter{result.IsValid()}
}
