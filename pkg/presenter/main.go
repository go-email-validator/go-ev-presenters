package presenter

import (
	"errors"
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"time"
)

type MultiplePresenter interface {
	SingleValidation(email string, name preparer.Name, opts map[ev.ValidatorName]interface{}) (interface{}, error)
}

func NewMultiplePresenter(presenters map[preparer.Name]Interface) MultiplePresenter {
	return multiplePresenter{presenters: presenters}
}

type multiplePresenter struct {
	presenters map[preparer.Name]Interface
}

func (p multiplePresenter) SingleValidation(email string, name preparer.Name, opts map[ev.ValidatorName]interface{}) (interface{}, error) {
	prep, ok := p.presenters[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("preparer with name \"%s\" does not exist", name))
	}

	return prep.SingleValidation(email, opts)
}

type GetEmail func(email string) evmail.Address

type Interface interface {
	SingleValidation(email string, opts map[ev.ValidatorName]interface{}) (interface{}, error)
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
func (p presenter) SingleValidation(email string, opts map[ev.ValidatorName]interface{}) (interface{}, error) {
	address := p.getEmail(email)

	start := time.Now()
	validationResult := p.validator.Validate(ev.NewInputFromMap(address, opts))
	elapsed := time.Since(start)

	return p.preparer.Prepare(address, validationResult, preparer.NewOptions(elapsed)), nil
}
