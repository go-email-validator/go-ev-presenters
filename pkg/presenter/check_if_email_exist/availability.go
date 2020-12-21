package check_if_email_exist

import "github.com/go-email-validator/go-ev-presenters/pkg/presenter/common"

type Availability string

func (a Availability) String() string {
	return string(a)
}

const (
	Risky   Availability = "risky"
	Invalid Availability = "invalid"
	Safe    Availability = "safe"
	Unknown Availability = "unknown"
)

func CalculateAvailability(depPresenter DepPresenter) Availability {
	if depPresenter.SMTP != common.FalseSMTPPresenter {
		if depPresenter.Misc.IsDisposable ||
			depPresenter.Misc.IsRoleAccount ||
			depPresenter.SMTP.IsCatchAll ||
			depPresenter.SMTP.HasFullInbox {
			return Risky
		}

		if !depPresenter.SMTP.IsDeliverable ||
			!depPresenter.SMTP.CanConnectSmtp ||
			depPresenter.SMTP.IsDisabled {
			return Invalid
		}

		return Safe
	}
	return Unknown
}
