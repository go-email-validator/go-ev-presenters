package _struct

import "github.com/go-email-validator/go-ev-presenters/pkg/presentation/prompt_email_verification_api"

// see prompt_email_verification_api/dep_functional_test.go
type DepPresentationTest struct {
	Email string
	Dep   prompt_email_verification_api.DepPresentation
}
