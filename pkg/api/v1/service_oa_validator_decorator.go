package v1

import (
	openapi "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/go"
	"github.com/go-email-validator/go-ev-presenters/pkg/log"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func NeApiOAValidatorDecorator(api openapi.EmailValidationApiRouter, openApiValidatorFactory openapi.ValidatorFactory, strictResponse bool) openapi.EmailValidationApiRouter {
	return &apiOAValidatorDecorator{
		api:            api,
		factory:        openApiValidatorFactory,
		strictResponse: strictResponse,
	}
}

type apiOAValidatorDecorator struct {
	api            openapi.EmailValidationApiRouter
	factory        openapi.ValidatorFactory
	strictResponse bool
}

func (a *apiOAValidatorDecorator) decorate(c *fiber.Ctx, fn func() error) error {
	openApiValidator := a.factory.Validator(c)
	if err := openApiValidator.ValidateRequest(c); err != nil {
		return ResponseError(c, err)
	}

	if err := fn(); err != nil {
		return err
	}

	if err := openApiValidator.ValidateResponse(c); err != nil {
		if a.strictResponse {
			return ResponseError(c, err)
		}
		log.Logger().Error("invalid response",
			zap.String("Method", c.Method()),
			zap.String("URL", c.BaseURL()),
			zap.ByteString("Body", c.Body()),
		)
	}

	return nil
}

func (a *apiOAValidatorDecorator) EmailValidationSingleValidationGet(c *fiber.Ctx) error {
	return a.decorate(c, func() error {
		return a.api.EmailValidationSingleValidationGet(c)
	})
}

func (a *apiOAValidatorDecorator) EmailValidationSingleValidationPost(c *fiber.Ctx) error {
	return a.decorate(c, func() error {
		return a.api.EmailValidationSingleValidationPost(c)
	})
}
