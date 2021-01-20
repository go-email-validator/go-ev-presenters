package v1

import "github.com/go-email-validator/go-email-validator/pkg/ev"

type mockPresenter struct {
	ret map[string]interface{}
}

func (m *mockPresenter) SingleValidation(email string, _ map[ev.ValidatorName]interface{}) (interface{}, error) {
	return m.ret[email], nil
}
