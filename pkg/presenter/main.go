package presenter

import (
	"errors"
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/prompt_email_verification_api"
	"time"
)

func NewMultiplePresentersDefault() MultiplePresenter {
	return MultiplePresenter{map[preparer.Name]Presenter{
		check_if_email_exist.Name: NewPresenter(
			ev_email.EmailFromString,
			check_if_email_exist.NewDepValidator(),
			check_if_email_exist.NewDepPreparerDefault(),
		),
		mailboxvalidator.Name: NewPresenter(
			mailboxvalidator.EmailFromString,
			mailboxvalidator.NewDepValidator(),
			mailboxvalidator.NewDepPreparerForViewDefault(),
		),
		prompt_email_verification_api.Name: NewPresenter(
			ev_email.EmailFromString,
			ev.NewDepBuilder(nil).Build(),
			prompt_email_verification_api.NewDepPreparerDefault(),
		),
	}}
}

type MultiplePresenter struct {
	presenters map[preparer.Name]Presenter
}

func (p MultiplePresenter) SingleValidation(email string, name preparer.Name) (interface{}, error) {
	prep, ok := p.presenters[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("preparer with name %s does not exist", name))
	}

	return prep.SingleValidation(email)
}

type GetEmail func(email string) ev_email.EmailAddress

func NewPresenter(getEmail GetEmail, validator ev.Validator, preparer preparer.Interface) Presenter {
	return Presenter{
		getEmail:  getEmail,
		validator: validator,
		preparer:  preparer,
	}
}

type Presenter struct {
	getEmail  func(email string) ev_email.EmailAddress
	validator ev.Validator
	preparer  preparer.Interface
}

func (p Presenter) SingleValidation(email string) (interface{}, error) {
	address := p.getEmail(email)

	start := time.Now()
	validationResult := p.validator.Validate(address)
	elapsed := time.Since(start)

	return p.preparer.Prepare(address, validationResult, preparer.NewOptions(elapsed)), nil
}
