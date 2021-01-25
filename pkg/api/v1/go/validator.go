package openapi

import (
	"bytes"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-email-validator/go-ev-presenters/statik"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Validator interface {
	Validate(c *fiber.Ctx) error
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
	httpReq.RemoteAddr = c.Context().RemoteIP().String()

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
