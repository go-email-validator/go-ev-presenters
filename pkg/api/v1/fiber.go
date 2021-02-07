package v1

import (
	"github.com/go-email-validator/go-ev-presenters/pkg/api/v1/go"
	"github.com/go-email-validator/go-ev-presenters/pkg/log"
	"github.com/gofiber/fiber/v2"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
)

type FiberFactory func(service openapi.EmailValidationApiRouter, opts Options) *fiber.App

func DefaultFiberFactory(service openapi.EmailValidationApiRouter, opts Options) *fiber.App {
	app := fiber.New(opts.Fiber)
	// Show error openapi format
	app.Use(func(c *fiber.Ctx) error {
		chainErr := c.Next()

		if chainErr != nil {
			return ResponseError(c, chainErr)
		}

		return chainErr
	})
	// logger
	app.Use(func(c *fiber.Ctx) error {
		chainErr := c.Next()

		// TODO Add formatting and fields
		if chainErr != nil {
			log.Logger().Error(chainErr.Error())
		}

		return chainErr
	})
	app.Use(fiberrecover.New())

	app.Post("/v1/validation/single", service.EmailValidationSingleValidationPost)
	app.Get("/v1/validation/single/:email", service.EmailValidationSingleValidationGet)

	return app
}
