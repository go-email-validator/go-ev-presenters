package main

type mockPresenter struct {
	ret map[string]interface{}
}

func (m *mockPresenter) SingleValidation(email string) (interface{}, error) {
	return m.ret[email], nil
}
