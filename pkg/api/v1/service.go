package v1

import (
	"encoding/json"
	openapi "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/go"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/url"
)

func Error(c *fiber.Ctx, err error) error {
	return c.JSON(openapi.UnexpectedError{Message: err.Error()})
}

func DefaultUnmarshal(c *fiber.Ctx, data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		c.Status(http.StatusBadRequest)
		return Error(c, err)
	}
	return nil
}

type unmarshal func(c *fiber.Ctx, data []byte, v interface{}) error

type EmailValidationApiControllerDTO struct {
	Presenter        presenter.MultiplePresenter
	Matching         map[openapi.ResultType]preparer.Name
	OpenApiValidator openapi.Validator
	JsonUnmarshal    unmarshal
}

func NewEmailValidationApiController(dto EmailValidationApiControllerDTO) openapi.EmailValidationApiRouter {
	if dto.JsonUnmarshal == nil {
		dto.JsonUnmarshal = DefaultUnmarshal
	}

	return &emailValidationApiController{
		presenter:        dto.Presenter,
		matching:         dto.Matching,
		openApiValidator: dto.OpenApiValidator,
		unmarshal:        dto.JsonUnmarshal,
	}
}

// A emailValidationApiController binds http requests to an api service and writes the service results to the http response
type emailValidationApiController struct {
	presenter        presenter.MultiplePresenter
	matching         map[openapi.ResultType]preparer.Name
	openApiValidator openapi.Validator
	unmarshal        unmarshal
}

func (e *emailValidationApiController) EmailValidationSingleValidationPost(c *fiber.Ctx) error {
	if err := e.openApiValidator.Validate(c); err != nil {
		return Error(c, err)
	}

	body := &openapi.EmailRequest{}
	if err := e.unmarshal(c, c.Request().Body(), &body); err != nil {
		return nil
	}

	if body.ResultType == "" {
		body.ResultType = openapi.CIEE
	}
	result, err := e.presenter.SingleValidation(body.Email, e.matching[body.ResultType])
	if err != nil {
		return Error(c, err)
	}

	return c.JSON(result)
}

var defaultResultType = string(openapi.CIEE)

func (e *emailValidationApiController) EmailValidationSingleValidationGet(c *fiber.Ctx) error {
	if err := e.openApiValidator.Validate(c); err != nil {
		return Error(c, err)
	}

	email, err := url.QueryUnescape(c.Params("email"))
	if err != nil {
		return Error(c, err)
	}
	resultType := openapi.ResultType(c.Query("result_type", defaultResultType))

	result, err := e.presenter.SingleValidation(email, e.matching[resultType])
	if err != nil {
		return Error(c, err)
	}

	return c.JSON(result)
}
