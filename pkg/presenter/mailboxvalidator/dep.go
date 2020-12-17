package mailboxvalidator

const (
	Name = "MailBoxValidator"

	MissingParameter    = 100
	ApiKeyNotFound      = MissingParameter + iota
	ApiKeyDisabled      = MissingParameter + iota
	ApiKeyExpired       = MissingParameter + iota
	InsufficientCredits = MissingParameter + iota
	UnknownError        = MissingParameter + iota
)

func NewDepPreparer() DepPresenter {
	panic("not implemented")
}

type DepPresenter struct {
	EmailAddress          string `json:"email_address"`
	Domain                string `json:"domain"`
	IsFree                string `json:"is_free"`
	IsSyntax              string `json:"is_syntax"`
	IsDomain              string `json:"is_domain"`
	IsSmtp                string `json:"is_smtp"`
	IsVerified            string `json:"is_verified"`
	IsServerDown          string `json:"is_server_down"`
	IsGreylisted          string `json:"is_greylisted"`
	IsDisposable          string `json:"is_disposable"`
	IsSuppressed          string `json:"is_suppressed"`
	IsRole                string `json:"is_role"`
	IsHighRisk            string `json:"is_high_risk"`
	IsCatchall            string `json:"is_catchall"`
	MailboxvalidatorScore string `json:"mailboxvalidator_score"`
	TimeTaken             string `json:"time_taken"`
	Status                string `json:"status"`
	CreditsAvailable      uint32 `json:"credits_available"`
	ErrorCode             string `json:"error_code"`
	ErrorMessage          string `json:"error_message"`
}
