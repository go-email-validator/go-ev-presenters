package presenter

import (
	"errors"
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/contains"
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
			ev.NewDepBuilder(nil).Build(),
			check_if_email_exist.NewDepPreparerDefault(),
		),
		mailboxvalidator.Name: NewPresenter(
			ev.NewDepBuilder(nil).Set(
				ev.BlackListEmailsValidatorName,
				ev.NewBlackListEmailsValidator(contains.NewSet(hashset.New(
					"example@example.com", "localhost@localhost",
				))),
			).Set(
				ev.BanWordsUsernameValidatorName,
				ev.NewBanWordsUsername(contains.NewInStringsFromArray([]string{"test"})),
			).Set(
				ev.FreeValidatorName,
				ev.FreeDefaultValidator(),
			).Build(),
			mailboxvalidator.NewDepPreparerForViewDefault(),
		),
		prompt_email_verification_api.Name: NewPresenter(
			ev.NewDepBuilder(nil).Build(),
			prompt_email_verification_api.NewDepPreparerDefault(),
		),
	}}
}

type MultiplePresenter struct {
	presenters map[preparer.Name]Presenter
}

func (p MultiplePresenter) SingleValidation(email ev_email.EmailAddress, name preparer.Name) (interface{}, error) {
	prep, ok := p.presenters[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("preparer with name %s does not exist", name))
	}

	return prep.SingleValidation(email)
}

func NewPresenter(validator ev.Validator, preparer preparer.Interface) Presenter {
	return Presenter{
		validator: validator,
		preparer:  preparer,
	}
}

type Presenter struct {
	validator ev.Validator
	preparer  preparer.Interface
}

func (p Presenter) SingleValidation(email ev_email.EmailAddress) (interface{}, error) {
	start := time.Now()
	validationResult := p.validator.Validate(email)
	elapsed := time.Since(start)

	return p.preparer.Prepare(email, validationResult, preparer.NewOptions(elapsed)), nil
}
