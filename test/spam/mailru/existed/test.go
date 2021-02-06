package main

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-ev-presenters/test/spam/common"
	"github.com/go-email-validator/go-ev-presenters/test/spam/mailru/utils"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	common.EmailsTest(
		utils.EmailAddresses[5:6],
		6000,
		ev.NewResult(true, nil, nil, ev.SMTPValidatorName),
	)
}
