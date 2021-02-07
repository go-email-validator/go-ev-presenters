package openapi

import (
	"bytes"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-email-validator/go-ev-presenters/statik"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

type ValidatorFactory interface {
	Validator(c *fiber.Ctx) Validator
}

type ValidatorFactoryFunc func(router *openapi3filter.Router, options *openapi3filter.Options, input *openapi3filter.RequestValidationInput) Validator

type Validator interface {
	ValidateRequest(c *fiber.Ctx) error
	ValidateResponse(c *fiber.Ctx) error
}

func RouterFromPath(path string) *openapi3filter.Router {
	openApiData, err := statik.ReadFile(path)
	if err != nil {
		panic(err)
	}
	router := openapi3.NewSwaggerLoader()
	swagger, err := router.LoadSwaggerFromData(openApiData)
	if err != nil {
		panic(err)
	}

	return openapi3filter.NewRouter().WithSwagger(swagger)
}

func NewValidatorFactory(router *openapi3filter.Router, options *openapi3filter.Options, validatorFactoryFunc ValidatorFactoryFunc) ValidatorFactory {
	if validatorFactoryFunc == nil {
		validatorFactoryFunc = NewValidator
	}

	return &validatorFactory{
		router:               router,
		options:              options,
		validatorFactoryFunc: validatorFactoryFunc,
	}
}

type validatorFactory struct {
	router               *openapi3filter.Router
	options              *openapi3filter.Options
	validatorFactoryFunc ValidatorFactoryFunc
}

func (v *validatorFactory) Validator(c *fiber.Ctx) Validator {
	httpReq, _ := http.NewRequest(c.Method(), c.Path(), bytes.NewReader(c.Body()))
	httpReq.RemoteAddr = c.Context().RemoteIP().String()

	c.Request().Header.VisitAll(func(k, v []byte) {
		sk := string(k)
		sv := string(v)
		switch sk {
		case "Transfer-Encoding":
			httpReq.TransferEncoding = append(httpReq.TransferEncoding, sv)
		case fiber.HeaderXForwardedFor:
			httpReq.Header.Set(fiber.HeaderXForwardedFor, strings.Join(c.IPs(), ", "))
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

	return v.validatorFactoryFunc(v.router, v.options, requestValidationInput)
}

func NewValidator(router *openapi3filter.Router, options *openapi3filter.Options, input *openapi3filter.RequestValidationInput) Validator {
	return &validator{
		router:  router,
		options: options,
		input:   input,
	}
}

type validator struct {
	router  *openapi3filter.Router
	options *openapi3filter.Options
	input   *openapi3filter.RequestValidationInput
}

func (v *validator) ValidateRequest(c *fiber.Ctx) error {
	return openapi3filter.ValidateRequest(c.Context(), v.input)
}

func (v *validator) ValidateResponse(c *fiber.Ctx) error {
	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: v.input,
		Status:                 c.Context().Response.StatusCode(),
		Header: http.Header{
			"Content-Type": []string{
				string(c.Context().Response.Header.ContentType()),
			},
		},
	}
	if body := c.Context().Response.Body(); body != nil {
		responseValidationInput.SetBodyBytes(body)
	}

	return openapi3filter.ValidateResponse(c.Context(), responseValidationInput)
}
