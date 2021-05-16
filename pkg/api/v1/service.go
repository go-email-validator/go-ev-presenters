package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp"
	"github.com/go-email-validator/go-email-validator/pkg/ev/utils"
	"github.com/go-email-validator/go-ev-presenters/pkg/api/v1/go"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/converter"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/url"
)

func ResponseError(c *fiber.Ctx, err error) error {
	return c.JSON(openapi.UnexpectedError{Message: err.Error()})
}

func DefaultUnmarshal(c *fiber.Ctx, data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		c.Status(http.StatusBadRequest)
		return ResponseError(c, err)
	}
	return nil
}

type unmarshal func(c *fiber.Ctx, data []byte, v interface{}) error

type EmailValidationApiControllerDTO struct {
	Presenter     presentation.MultiplePresenter
	Matching      map[openapi.ResultType]converter.Name
	JsonUnmarshal unmarshal
	// MatchingResponse needs to juxtapose with openapi.EmailResponse
	MatchingResponse map[converter.Name]string
}

func NewEmailValidationApiController(dto EmailValidationApiControllerDTO) openapi.EmailValidationApiRouter {
	if dto.JsonUnmarshal == nil {
		dto.JsonUnmarshal = DefaultUnmarshal
	}

	return &emailValidationApiController{
		presenter:        dto.Presenter,
		matching:         dto.Matching,
		unmarshal:        dto.JsonUnmarshal,
		matchingResponse: dto.MatchingResponse,
	}
}

// A emailValidationApiController binds http requests to an api service and writes the service results to the http response
type emailValidationApiController struct {
	presenter        presentation.MultiplePresenter
	matching         map[openapi.ResultType]converter.Name
	matchingResponse map[converter.Name]string
	unmarshal        unmarshal
}

var smtpDefaultOpts = evsmtp.DefaultOptions()
var gravatarDefaultOpts = ev.DefaultGravatarOptions()

func (e *emailValidationApiController) preparePresenter(name converter.Name, presenter interface{}) interface{} {
	if key, ok := e.matchingResponse[name]; ok {
		return map[string]interface{}{
			key: presenter,
		}
	}

	panic(fmt.Sprintf("Value does not exist for key \"%s\"", name))
}

func (e *emailValidationApiController) EmailValidationSingleValidationPost(c *fiber.Ctx) error {
	body := &openapi.EmailRequest{}
	if err := e.unmarshal(c, c.Request().Body(), &body); err != nil {
		return nil
	}

	if body.ResultType == "" {
		body.ResultType = openapi.CIEE
	}

	opts := map[ev.ValidatorName]interface{}{
		ev.SMTPValidatorName: evsmtp.NewOptions(evsmtp.OptionsDTO{
			EmailFrom: evmail.EmptyEmail(
				evmail.FromString(body.Smtp.EmailFrom),
				smtpDefaultOpts.EmailFrom(),
			),
			HelloName: utils.DefaultString(
				body.Smtp.HelloName,
				smtpDefaultOpts.HelloName(),
			),
			Proxy: utils.DefaultString(
				body.Smtp.Proxy,
				smtpDefaultOpts.Proxy(),
			),
			TimeoutCon: utils.DefaultDuration(
				OnlyPositiveDuration(SecondDuration(float64(body.Smtp.TimeoutConnection))),
				gravatarDefaultOpts.Timeout(),
			),
			TimeoutResp: utils.DefaultDuration(
				OnlyPositiveDuration(SecondDuration(float64(body.Smtp.TimeoutResponse))),
				gravatarDefaultOpts.Timeout(),
			),
			Port: utils.DefaultInt(
				OnlyPositiveInt(int(body.Smtp.Port)),
				smtpDefaultOpts.Port(),
			),
		}),
		/*ev.GravatarValidatorName: ev.NewGravatarOptions(ev.GravatarOptionsDTO{
			Timeout: utils.DefaultDuration(
				OnlyPositiveDuration(SecondDuration(float64(body.Gravatar.Timeout))),
				gravatarDefaultOpts.Timeout(),
			),
		}),*/
	}

	preparerName := e.matching[body.ResultType]
	result, err := e.presenter.Validate(body.Email, preparerName, opts)
	if err != nil {
		return ResponseError(c, err)
	}
	return c.JSON(e.preparePresenter(preparerName, result))
}

var defaultResultType = string(openapi.CIEE)

func (e *emailValidationApiController) EmailValidationSingleValidationGet(c *fiber.Ctx) error {
	email, err := url.QueryUnescape(c.Params("email"))
	if err != nil {
		return ResponseError(c, err)
	}
	resultType := openapi.ResultType(c.Query("result_type", defaultResultType))

	preparerName := e.matching[resultType]
	result, err := e.presenter.Validate(email, preparerName, nil)
	if err != nil {
		return ResponseError(c, err)
	}

	return c.JSON(e.preparePresenter(preparerName, result))
}
