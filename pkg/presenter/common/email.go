package common

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"strings"
)

func NewEmailAddress(username, domain string, at *string) ev_email.EmailAddress {
	return emailAddress{
		username: strings.ToLower(username),
		at:       at,
		domain:   strings.ToLower(domain),
	}
}

type emailAddress struct {
	username string
	at       *string
	domain   string
}

func (e emailAddress) Username() string {
	return e.username
}

func (e emailAddress) Domain() string {
	return e.domain
}

func (e emailAddress) String() string {
	if e.at == nil {
		return e.Username() + ev_email.AT + e.Domain()
	}

	return e.Username() + *e.at + e.Domain()
}
