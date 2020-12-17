package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
)

const Name preparer.Name = "Name"

type miscPresenter struct {
	disposablePresenter
	rolePresenter
}

type DepPresenter struct {
	Input       string          `json:"input"`
	IsReachable availability    `json:"is_reachable"`
	Misc        miscPresenter   `json:"misc"`
	MX          mxPresenter     `json:"mx"`
	SMTP        smtpPresenter   `json:"smtp"`
	Syntax      syntaxPresenter `json:"syntax"`
}

func NewDepPreparer() DepPreparer {
	return DepPreparer{
		preparer.NewMultiplePreparer(preparer.MapPreparers{
			ev.RoleValidatorName:       rolePreparer{},
			ev.DisposableValidatorName: disposablePreparer{},
			ev.MXValidatorName:         mxPreparer{},
			ev.SMTPValidatorName:       SMTPPreparer{},
			ev.SyntaxValidatorName:     SyntaxPreparer{},
		}),
		calculateAvailability,
	}
}

type DepPreparer struct {
	preparer              preparer.MultiplePreparer
	calculateAvailability func(depPresenter DepPresenter) availability
}

func (s DepPreparer) CanPrepare(_ email.EmailAddressInterface, result ev.ValidationResultInterface) bool {
	return result.ValidatorName() == ev.DepValidatorName
}

func (s DepPreparer) Prepare(email email.EmailAddressInterface, result ev.ValidationResultInterface) interface{} {
	depPresenter := DepPresenter{
		Input: email.String(),
		Misc:  miscPresenter{},
	}

	for _, validatorResult := range result.(ev.DepValidatorResultInterface).GetResults() {
		if !s.preparer.CanPrepare(email, validatorResult) {
			continue
		}

		switch v := s.preparer.Prepare(email, validatorResult).(type) {
		case rolePresenter:
			depPresenter.Misc.rolePresenter = v
		case disposablePresenter:
			depPresenter.Misc.disposablePresenter = v
		case mxPresenter:
			depPresenter.MX = v
		case smtpPresenter:
			depPresenter.SMTP = v
		case syntaxPresenter:
			depPresenter.Syntax = v
		}
	}
	depPresenter.IsReachable = s.calculateAvailability(depPresenter)

	return depPresenter
}
