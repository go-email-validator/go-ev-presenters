package v1

import (
	"context"
	"errors"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/getkin/kin-openapi/openapi3filter"
)

var (
	ErrAuthApiKey = errors.New("apiKey is incorrect")
	ErrAuthIP     = errors.New("IP \"%s\" is not acceptable")
)

type AuthFunc func(c context.Context, input *openapi3filter.AuthenticationInput) error

func AuthFuncComplex(opts AuthOptions) AuthFunc {
	return AuthenticationFuncIPs(
		opts,
		AuthenticationFuncHeaders(opts, nil),
	)
}

func AuthenticationFuncIPs(opts AuthOptions, next AuthFunc) AuthFunc {
	if len(opts.IPs) == 0 {
		return next
	}

	IIPs := make([]interface{}, len(opts.IPs))
	for i, IP := range opts.IPs {
		IIPs[i] = IP
	}

	IPs := hashset.New(IIPs...)
	IIPs = nil

	return func(c context.Context, input *openapi3filter.AuthenticationInput) error {
		if !IPs.Contains(input.RequestValidationInput.Request.RemoteAddr) {
			return input.NewError(ErrAuthApiKey)
		}

		if next != nil {
			return next(c, input)
		}

		return nil
	}
}

func AuthenticationFuncHeaders(opts AuthOptions, next AuthFunc) AuthFunc {
	if len(opts.Headers) == 0 {
		return next
	}

	return func(c context.Context, input *openapi3filter.AuthenticationInput) error {
		for key, value := range opts.Headers {
			if input.RequestValidationInput.Request.Header.Get(key) != value {
				return input.NewError(ErrAuthIP)
			}
		}

		if next != nil {
			return next(c, input)
		}

		return nil
	}
}
