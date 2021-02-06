package main

import (
	"errors"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-ev-presenters/test/spam/common"
	"runtime"
	"strconv"
)

func main() {
	runtime.GOMAXPROCS(1)

	randomEmails := make([]evmail.Address, 0)
	for i := 0; i < 2000; i++ {
		randomEmails = append(randomEmails, evmail.FromString("some.non.existed.email"+strconv.Itoa(i)+"@gmail.com"))
	}

	common.EmailsTest(
		randomEmails,
		0,
		ev.NewResult(false, []error{errors.New("")}, []error{errors.New("")}, ev.SMTPValidatorName),
	)
}
