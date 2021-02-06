package main

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-ev-presenters/test/spam/common"
	"runtime"
	"strconv"
)

func main() {
	runtime.GOMAXPROCS(1)
	randomEmails := make([]evmail.Address, 0)
	for i := 0; i < 200; i++ {
		randomEmails = append(randomEmails, evmail.FromString("some.non.existed.email"+strconv.Itoa(i)+"@mail.ru"))
	}

	common.EmailsTest(
		randomEmails,
		0,
		ev.NewResult(true, nil, nil, ev.SMTPValidatorName),
	)
}
