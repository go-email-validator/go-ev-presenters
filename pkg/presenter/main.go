package presenter

import (
	"errors"
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
)

func NewPresenter() Presenter {
	return Presenter{
		validator: ev.NewDepBuilder(nil).Build(),
		preparers: map[preparer.Name]preparer.Interface{
			check_if_email_exist.Name: check_if_email_exist.NewDepPreparer(),
		},
	}
}

type Presenter struct {
	validator ev.ValidatorInterface
	preparers map[preparer.Name]preparer.Interface
}

func (p Presenter) SingleValidation(email ev_email.EmailAddressInterface, name preparer.Name) (interface{}, error) {
	prep, ok := p.preparers[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("preparer with name %s does not exist", name))
	}

	return prep.Prepare(email, p.validator.Validate(email)), nil
}
