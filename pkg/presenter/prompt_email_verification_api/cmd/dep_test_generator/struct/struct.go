package _struct

import "github.com/go-email-validator/go-ev-presenters/pkg/presenter/prompt_email_verification_api"

// see /pkg/presenter/prompt_email_verification_api/dep_functional_test.go
type DepPresenterTest struct {
	Email string
	Dep   prompt_email_verification_api.DepPresenter
}
