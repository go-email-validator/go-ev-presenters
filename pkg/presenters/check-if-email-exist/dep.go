package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenters/presenter"
)

type MiscPresenter struct {
	DisposablePresenter
	RolePresenter
}

type DepPresenter struct {
	Input       string          `json:"input"`
	IsReachable Availability    `json:"is_reachable"`
	Misc        MiscPresenter   `json:"misc"`
	MX          MXPresenter     `json:"mx"`
	SMTP        SMTPPresenter   `json:"smtp"`
	Syntax      SyntaxPresenter `json:"syntax"`
}

func NewDepProcessor() DepProcessor {
	return DepProcessor{
		presenter.NewMultiplePreparer(presenter.MapPreparers{
			ev.RoleValidatorName:       RoleProcessor{},
			ev.DisposableValidatorName: DisposableProcessor{},
			ev.MXValidatorName:         MXProcessor{},
			ev.SMTPValidatorName:       SMTPProcessor{},
			ev.SyntaxValidatorName:     SyntaxProcessor{},
		}),
		CalculateAvailability,
	}
}

type DepProcessor struct {
	processor             presenter.MultiplePreparer
	calculateAvailability func(depPresenter DepPresenter) Availability
}

func (s DepProcessor) CanProcess(_ email.EmailAddressInterface, result ev.ValidationResultInterface) bool {
	return result.ValidatorName() == ev.DepValidatorName
}

func (s DepProcessor) Process(email email.EmailAddressInterface, result ev.ValidationResultInterface) interface{} {
	depPresenter := DepPresenter{
		Input: email.String(),
		Misc:  MiscPresenter{},
	}

	for _, result := range result.(ev.DepValidatorResultInterface).GetResults() {
		if !s.processor.CanProcess(email, result) {
			continue
		}

		switch v := s.processor.Process(email, result).(type) {
		case RolePresenter:
			depPresenter.Misc.RolePresenter = v
		case DisposablePresenter:
			depPresenter.Misc.DisposablePresenter = v
		case MXPresenter:
			depPresenter.MX = v
		case SMTPPresenter:
			depPresenter.SMTP = v
		case SyntaxPresenter:
			depPresenter.Syntax = v
		}
	}
	depPresenter.IsReachable = s.calculateAvailability(depPresenter)

	return depPresenter
}
