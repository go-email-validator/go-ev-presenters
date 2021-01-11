package openapi

import (
	"bytes"
	"context"
	"errors"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var (
	ErrAuthApiKey = errors.New("apiKey is incorrect")
)

type Validator interface {
	Validate(c *fiber.Ctx) error
}

func RouterFromPath(path string) *openapi3filter.Router {
	return openapi3filter.NewRouter().WithSwaggerFromFile(path)
}

func AuthenticationFuncWithKey(apiKey string) func(c context.Context, input *openapi3filter.AuthenticationInput) error {
	return func(c context.Context, input *openapi3filter.AuthenticationInput) error {
		if input.RequestValidationInput.Request.Header.Get(input.SecurityScheme.Name) != apiKey {
			return input.NewError(ErrAuthApiKey)
		}
		return nil
	}
}

func NewValidator(router *openapi3filter.Router, options *openapi3filter.Options) Validator {
	return &validator{
		router:  router,
		options: options,
	}
}

type validator struct {
	router  *openapi3filter.Router
	options *openapi3filter.Options
}

func (v *validator) Validate(c *fiber.Ctx) error {
	httpReq, _ := http.NewRequest(c.Method(), c.Path(), bytes.NewReader(c.Body()))

	c.Request().Header.VisitAll(func(k, v []byte) {
		sk := string(k)
		sv := string(v)
		switch sk {
		case "Transfer-Encoding":
			httpReq.TransferEncoding = append(httpReq.TransferEncoding, sv)
		default:
			httpReq.Header.Set(sk, sv)
		}
	})

	route, pathParams, _ := v.router.FindRoute(c.Method(), httpReq.URL)
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    httpReq,
		PathParams: pathParams,
		Route:      route,
		Options:    v.options,
	}

	return openapi3filter.ValidateRequest(c.Context(), requestValidationInput)
}
