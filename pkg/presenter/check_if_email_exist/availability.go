package check_if_email_exist

type availability string

func (a availability) String() string {
	return string(a)
}

const (
	Risky   availability = "risky"
	Invalid availability = "invalid"
	Safe    availability = "safe"
	Unknown availability = "unknown"
)

func calculateAvailability(depPresenter DepPresenter) availability {
	if depPresenter.SMTP != FalseSMTPPresenter {
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
	} else {
		return Unknown
	}
}
