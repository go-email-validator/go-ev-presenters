package mailboxvalidator

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"time"
)

const (
	Name preparer.Name = "MailBoxValidator"

	MissingParameter        int32 = 100
	MissingParameterMessage       = "Missing parameter."
	ApiKeyNotFound                = MissingParameter + iota
	ApiKeyDisabled                = MissingParameter + iota
	ApiKeyExpired                 = MissingParameter + iota
	InsufficientCredits           = MissingParameter + iota
	UnknownError                  = MissingParameter + iota
	UnknownErrorMessage           = "Unknown error."
)

type DepPresenter struct {
	EmailAddress          string        `json:"email_address"`
	Domain                string        `json:"domain"`
	IsFree                bool          `json:"is_free"`
	IsSyntax              bool          `json:"is_syntax"`
	IsDomain              bool          `json:"is_domain"`
	IsSmtp                bool          `json:"is_smtp"`
	IsVerified            bool          `json:"is_verified"`
	IsServerDown          bool          `json:"is_server_down"`
	IsGreylisted          bool          `json:"is_greylisted"`
	IsDisposable          bool          `json:"is_disposable"`
	IsSuppressed          bool          `json:"is_suppressed"`
	IsRole                bool          `json:"is_role"`
	IsHighRisk            bool          `json:"is_high_risk"`
	IsCatchall            bool          `json:"is_catchall"`
	MailboxvalidatorScore float64       `json:"mailboxvalidator_score"`
	TimeTaken             time.Duration `json:"time_taken"`
	Status                bool          `json:"status"`
	CreditsAvailable      uint32        `json:"credits_available"`
	ErrorCode             string        `json:"error_code"`
	ErrorMessage          string        `json:"error_message"`
}

type FuncCalculateScore func(presenter DepPresenter) float64

func NewDepPreparerDefault() DepPreparer {
	return NewDepPreparer(CalculateScore)
}

func NewDepPreparer(calculateScore FuncCalculateScore) DepPreparer {
	return DepPreparer{calculateScore}
}

type DepPreparer struct {
	calculateScore FuncCalculateScore
}

func (_ DepPreparer) CanPrepare(_ email.EmailAddress, result ev.ValidationResult, opts preparer.Options) bool {
	return opts.ExecutedTime() != 0 && result.ValidatorName() == ev.DepValidatorName
}

func (d DepPreparer) Prepare(email email.EmailAddress, resultInterface ev.ValidationResult, opts preparer.Options) (result interface{}) {
	defer func() {
		if r := recover(); r != nil {
			result = DepPresenter{
				ErrorCode:    string(MissingParameter),
				ErrorMessage: UnknownErrorMessage,
			}
		}
	}()
	var depPresenter DepPresenter
	if len(email.String()) == 0 {
		return DepPresenter{
			ErrorCode:    string(MissingParameter),
			ErrorMessage: MissingParameterMessage,
		}
	}

	depResult := resultInterface.(ev.DepValidationResult)
	validationResults := depResult.GetResults()

	smtpPresenter := common.SMTPPreparer{}.Prepare(email, validationResults[ev.SMTPValidatorName], nil).(common.SmtpPresenter)

	depPresenter = DepPresenter{
		EmailAddress:     email.String(),
		Domain:           email.Domain(),
		IsFree:           validationResults[ev.FreeValidatorName].IsValid(),
		IsSyntax:         validationResults[ev.SyntaxValidatorName].IsValid(),
		IsDomain:         validationResults[ev.MXValidatorName].IsValid(),
		IsSmtp:           smtpPresenter.CanConnectSmtp,
		IsVerified:       smtpPresenter.IsDeliverable,
		IsServerDown:     !smtpPresenter.CanConnectSmtp,
		IsGreylisted:     smtpPresenter.IsGreyListed,
		IsDisposable:     validationResults[ev.DisposableValidatorName].IsValid(),
		IsSuppressed:     !validationResults[ev.BlackListEmailsValidatorName].IsValid(), // TODO find more examples example@example.com
		IsRole:           validationResults[ev.RoleValidatorName].IsValid(),
		IsHighRisk:       !validationResults[ev.BanWordsUsernameValidatorName].IsValid(), // TODO find more words
		IsCatchall:       smtpPresenter.IsCatchAll,
		TimeTaken:        opts.ExecutedTime(),
		Status:           resultInterface.IsValid(), // valid or not use warning
		CreditsAvailable: ^uint32(0),
		ErrorCode:        "",
		ErrorMessage:     "",
	}

	depPresenter.MailboxvalidatorScore = d.calculateScore(depPresenter)
	return depPresenter
}
