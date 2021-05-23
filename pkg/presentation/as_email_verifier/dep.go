package as_email_verifier

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/converter"
	"regexp"
)

const (
	Name converter.Name = "AfterShipEmailVerifier"

	EmailRegexString = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
)

type DepPresentation struct {
	Email        string                `json:"email"`
	Reachable    Reachable             `json:"reachable"`
	Syntax       *SyntaxPresentation   `json:"syntax"`
	SMTP         *SmtpPresentation     `json:"smtp"`
	Gravatar     *GravatarPresentation `json:"gravatar"`
	Suggestion   string                `json:"suggestion"`
	Disposable   bool                  `json:"disposable"`
	RoleAccount  bool                  `json:"role_account"`
	Free         bool                  `json:"free"`
	HasMxRecords bool                  `json:"has_mx_records"`
}

type DepConverter struct {
	converter          converter.CompositeConverter
	calculateReachable FuncReachable
}

func NewDepConverterDefault() DepConverter {
	return NewDepConverter(
		converter.NewCompositeConverter(converter.MapConverters{
			ev.SMTPValidatorName:     converter.NewSMTPConverter(),
			ev.SyntaxValidatorName:   SyntaxConverter{},
			ev.GravatarValidatorName: GravatarConverter{},
		}),
		CalculateReachable,
	)
}

func NewDepConverter(converter converter.CompositeConverter, calculateReachable FuncReachable) DepConverter {
	return DepConverter{converter, calculateReachable}
}

func (DepConverter) Can(_ evmail.Address, result ev.ValidationResult, _ converter.Options) bool {
	return result.ValidatorName() == ev.DepValidatorName
}

func (s DepConverter) Convert(email evmail.Address, resultInterface ev.ValidationResult, opts converter.Options) interface{} {
	depResult := resultInterface.(ev.DepValidationResult)
	validationResults := depResult.GetResults()

	mxResult := validationResults[ev.MXValidatorName].(ev.MXValidationResult)

	depPresentation := DepPresentation{
		Email:        email.String(),
		Disposable:   validationResults[ev.DisposableValidatorName].IsValid(),
		RoleAccount:  validationResults[ev.RoleValidatorName].IsValid(),
		Free:         validationResults[ev.FreeValidatorName].IsValid(),
		HasMxRecords: len(mxResult.MX()) > 0,
		Suggestion:   "", // TODO need to resolve
	}

	for _, validatorResult := range depResult.GetResults() {
		if !s.converter.Can(email, validatorResult, opts) {
			continue
		}

		switch v := s.converter.Convert(email, validatorResult, opts).(type) {
		case *SyntaxPresentation:
			depPresentation.Syntax = v
		case converter.SmtpPresentation:
			depPresentation.SMTP = &SmtpPresentation{
				HostExists:  v.CanConnectSmtp,
				FullInbox:   v.HasFullInbox,
				CatchAll:    v.IsCatchAll,
				Deliverable: v.IsDeliverable,
				Disabled:    v.IsDisabled,
			}
		case *GravatarPresentation:
			depPresentation.Gravatar = v
		}
	}
	depPresentation.Reachable = s.calculateReachable(depPresentation)

	return depPresentation
}

func NewDepValidator(smtpValidator ev.Validator) ev.Validator {
	builder := ev.NewDepBuilder(nil)
	if smtpValidator == nil {
		smtpValidator = builder.Get(ev.SMTPValidatorName)
	}

	return ev.NewDepBuilder(nil).Set(
		ev.SyntaxValidatorName,
		ev.NewSyntaxRegexValidator(regexp.MustCompile(EmailRegexString)),
	).
		Set(ev.GravatarValidatorName, ev.NewGravatarValidator()).
		Set(ev.SMTPValidatorName, smtpValidator).Build()
}