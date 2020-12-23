package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
)

const Name preparer.Name = "CheckIfEmailExist"

type miscPresenter struct {
	disposablePresenter
	rolePresenter
}

// https://github.com/amaurymartiny/check-if-email-exists or https://reacher.email
type DepPresenter struct {
	Input       string               `json:"input"`
	IsReachable Availability         `json:"is_reachable"`
	Misc        miscPresenter        `json:"misc"`
	MX          mxPresenter          `json:"mx"`
	SMTP        common.SmtpPresenter `json:"smtp"`
	Syntax      syntaxPresenter      `json:"syntax"`
}

type FuncAvailability func(depPresenter DepPresenter) Availability

func NewDepPreparerDefault() DepPreparer {
	return NewDepPreparer(
		preparer.NewMultiplePreparer(preparer.MapPreparers{
			ev.RoleValidatorName:       rolePreparer{},
			ev.DisposableValidatorName: disposablePreparer{},
			ev.MXValidatorName:         mxPreparer{},
			ev.SMTPValidatorName:       common.SMTPPreparer{},
			ev.SyntaxValidatorName:     SyntaxPreparer{},
		}),
		CalculateAvailability,
	)
}

func NewDepPreparer(preparer preparer.MultiplePreparer, calculateAvailability FuncAvailability) DepPreparer {
	return DepPreparer{preparer, calculateAvailability}
}

type DepPreparer struct {
	preparer              preparer.MultiplePreparer
	calculateAvailability FuncAvailability
}

func (_ DepPreparer) CanPrepare(_ email.EmailAddress, result ev.ValidationResult, _ preparer.Options) bool {
	return result.ValidatorName() == ev.DepValidatorName
}

func (s DepPreparer) Prepare(email email.EmailAddress, result ev.ValidationResult, opts preparer.Options) interface{} {
	depPresenter := DepPresenter{
		Input: email.String(),
		Misc:  miscPresenter{},
	}

	for _, validatorResult := range result.(ev.DepValidationResult).GetResults() {
		if !s.preparer.CanPrepare(email, validatorResult, opts) {
			continue
		}

		switch v := s.preparer.Prepare(email, validatorResult, opts).(type) {
		case rolePresenter:
			depPresenter.Misc.rolePresenter = v
		case disposablePresenter:
			depPresenter.Misc.disposablePresenter = v
		case mxPresenter:
			depPresenter.MX = v
		case common.SmtpPresenter:
			depPresenter.SMTP = v
		case syntaxPresenter:
			depPresenter.Syntax = v
		}
	}
	depPresenter.IsReachable = s.calculateAvailability(depPresenter)

	return depPresenter
}
