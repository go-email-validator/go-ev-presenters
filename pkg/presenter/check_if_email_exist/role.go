package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
)

type rolePresenter struct {
	IsRoleAccount bool `json:"is_role_account"`
}

type rolePreparer struct{}

func (rolePreparer) CanPrepare(_ evmail.Address, result ev.ValidationResult, _ preparer.Options) bool {
	return result.ValidatorName() == ev.RoleValidatorName
}

func (rolePreparer) Prepare(_ evmail.Address, result ev.ValidationResult, _ preparer.Options) interface{} {
	return rolePresenter{!result.IsValid()}
}
