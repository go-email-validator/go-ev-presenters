package prompt_email_verification_api

import (
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common"
	"strings"
)

var emptyString = ""

func EmailFromString(email string) email.EmailAddress {
	firstPos := strings.IndexByte(email, '@')
	lastPos := strings.LastIndexByte(email, '@')

	if firstPos == -1 || len(email) < 3 || firstPos != lastPos {
		return common.NewEmailAddress("", "", &emptyString)
	}

	return common.NewEmailAddress(email[:firstPos], email[firstPos+1:], nil)
}
