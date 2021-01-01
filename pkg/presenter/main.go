package presenter

import (
	"errors"
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/prompt_email_verification_api"
	"time"
)

func NewMultiplePresentersDefault() MultiplePresenter {
	return NewMultiplePresenter(map[preparer.Name]Interface{
		check_if_email_exist.Name: NewPresenter(
			evmail.FromString,
			check_if_email_exist.NewDepValidator(),
			check_if_email_exist.NewDepPreparerDefault(),
		),
		mailboxvalidator.Name: NewPresenter(
			mailboxvalidator.EmailFromString,
			mailboxvalidator.NewDepValidator(),
			mailboxvalidator.NewDepPreparerForViewDefault(),
		),
		prompt_email_verification_api.Name: NewPresenter(
			prompt_email_verification_api.EmailFromString,
			prompt_email_verification_api.NewDepValidator(),
			prompt_email_verification_api.NewDepPreparerDefault(),
		),
	})
}

type MultiplePresenter interface {
	SingleValidation(email string, name preparer.Name) (interface{}, error)
}

func NewMultiplePresenter(presenters map[preparer.Name]Interface) MultiplePresenter {
	return multiplePresenter{presenters: presenters}
}

type multiplePresenter struct {
	presenters map[preparer.Name]Interface
}

func (p multiplePresenter) SingleValidation(email string, name preparer.Name) (interface{}, error) {
	prep, ok := p.presenters[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("preparer with name %s does not exist", name))
	}

	return prep.SingleValidation(email)
}

type GetEmail func(email string) evmail.Address

type Interface interface {
	SingleValidation(email string) (interface{}, error)
}

func NewPresenter(getEmail GetEmail, validator ev.Validator, preparer preparer.Interface) Interface {
	return presenter{
		getEmail:  getEmail,
		validator: validator,
		preparer:  preparer,
	}
}

type presenter struct {
	getEmail  func(email string) evmail.Address
	validator ev.Validator
	preparer  preparer.Interface
}

// TODO if error will be put, mockPresenter should return it
func (p presenter) SingleValidation(email string) (interface{}, error) {
	address := p.getEmail(email)

	start := time.Now()
	validationResult := p.validator.Validate(address)
	elapsed := time.Since(start)

	return p.preparer.Prepare(address, validationResult, preparer.NewOptions(elapsed)), nil
}
