package main

import (
	"errors"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-ev-presenters/test/spam/common"
	"github.com/go-email-validator/go-ev-presenters/test/spam/google/utils"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	common.EmailsTest(
		utils.EmailAddresses,
		common.MaxRequestPerOneEmail,
		ev.NewResult(true, nil, []error{errors.New("")}, ev.SMTPValidatorName),
	)
}
