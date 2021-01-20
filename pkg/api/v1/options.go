package v1

import (
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp"
	openapi "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/go"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/prompt_email_verification_api"
	"github.com/gofiber/fiber/v2"
)

const (
	HTTPDefaultHost = "0.0.0.0:8080"
	SwaggerPath     = "api/v1/openapiv3/ev.openapiv3.yaml"
)

func defaultInstance(opts Options) openapi.EmailValidationApiRouter {
	return NewEmailValidationApiController(EmailValidationApiControllerDTO{
		Presenter: getPresenter(evsmtp.CheckerDTO{
			SendMailFactory: evsmtp.NewSendMailFactory(evsmtp.H12IODial, nil),
			Options: evsmtp.NewOptions(evsmtp.OptionsDTO{
				HelloName: opts.Validator.HelloName,
				Proxy:     opts.Validator.SMTPProxy,
			}),
		}),
		Matching: map[openapi.ResultType]preparer.Name{
			openapi.CIEE:                          check_if_email_exist.Name,
			openapi.CHECK_IF_EMAIL_EXIST:          check_if_email_exist.Name,
			openapi.MBV:                           mailboxvalidator.Name,
			openapi.MAILBOXVALIDATOR:              mailboxvalidator.Name,
			openapi.MAIL_BOX_VALIDATOR:            mailboxvalidator.Name,
			openapi.PEVA:                          prompt_email_verification_api.Name,
			openapi.PROMPT_EMAIL_VERIFICATION_API: prompt_email_verification_api.Name,
		},
		OpenApiValidator: openapi.NewValidator(
			openapi.RouterFromPath(opts.HTTP.OpenApiPath),
			&openapi3filter.Options{
				AuthenticationFunc: openapi.AuthenticationFuncWithKey(opts.Auth.Key),
			},
		),
	})
}

var getPresenter = presenter.NewMultiplePresentersDefault

func NewOptions() Options {
	return Options{
		Validator: NewValidator(),
		HTTP:      NewHTTPOptions(),
		Auth:      NewAuthOptions(),
		Fiber:     fiber.Config{DisableStartupMessage: true},
	}
}

type Options struct {
	Validator Validator
	HTTP      HTTPOptions
	Auth      AuthOptions
	Fiber     fiber.Config
}

func NewValidator() Validator {
	return Validator{}
}

type Validator struct {
	SMTPProxy string
	HelloName string
}

func NewHTTPOptions() HTTPOptions {
	return HTTPOptions{
		Bind:        HTTPDefaultHost,
		OpenApiPath: SwaggerPath,
	}
}

type HTTPOptions struct {
	Bind        string
	OpenApiPath string
}

func NewAuthOptions() AuthOptions {
	return AuthOptions{}
}

type AuthOptions struct {
	Key string
}
