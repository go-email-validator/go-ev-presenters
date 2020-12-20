package presenter

import (
	"errors"
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"time"
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

	t := 1500 * time.Nanosecond
	fmt.Print(fmt.Sprint(t.Round(time.Microsecond).Seconds()))

	start := time.Now()
	validationResult := p.validator.Validate(email)
	elapsed := time.Since(start)

	return prep.Prepare(email, validationResult, preparer.Options{ExecutedTimeValue: elapsed}), nil
}
