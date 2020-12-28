package mailboxvalidator

import (
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common"
	"strings"
)

var emptyString = ""

func EmailFromString(email string) email.EmailAddress {
	pos := strings.LastIndexByte(email, '@')

	if pos == -1 || len(email) < 3 {
		return common.NewEmailAddress("", email, &emptyString)
	}

	return common.NewEmailAddress(email[:pos], email[pos+1:], nil)
}
