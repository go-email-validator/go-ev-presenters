package main

import (
	mockevsmtp "github.com/go-email-validator/go-email-validator/test/mock/ev/evsmtp"
	"testing"
)

func main() {
	t := &testing.T{}
	mockevsmtp.Server(t, mockevsmtp.SuccessServer, 0, "127.0.0.1:25", true)
}
