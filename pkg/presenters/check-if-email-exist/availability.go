package check_if_email_exist

type Availability string

const (
	Risky   Availability = "risky"
	Invalid Availability = "invalid"
	Safe    Availability = "safe"
	Unknown Availability = "unknown"
)

func CalculateAvailability(depPresenter DepPresenter) Availability {
	if depPresenter.SMTP == FalseSMTPPresenter {
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
