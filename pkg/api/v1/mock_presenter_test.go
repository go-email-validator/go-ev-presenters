package v1

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
)

type mockPresenter struct {
	ret map[string]interface{}
}

func (m *mockPresenter) Validate(email string, _ map[ev.ValidatorName]interface{}) (interface{}, error) {
	return m.ret[email], nil
}

func (m *mockPresenter) ValidateFromAddress(_ evmail.Address, _ map[ev.ValidatorName]interface{}) (interface{}, error) {
	panic("Method ValidateFromAddress is not realized")
}
