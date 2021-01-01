package mailboxvalidator

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/contains"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"time"
)

//go:generate go run cmd/dep_test_generator/gen.go

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

// https://www.mailboxvalidator.com/
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

func (DepPreparer) CanPrepare(_ evmail.Address, result ev.ValidationResult, opts preparer.Options) bool {
	return opts.ExecutedTime() != 0 && result.ValidatorName() == ev.DepValidatorName
}

func (d DepPreparer) Prepare(email evmail.Address, resultInterface ev.ValidationResult, opts preparer.Options) (result interface{}) {
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

	isFree := !validationResults[ev.FreeValidatorName].IsValid()
	isSyntax := validationResults[ev.SyntaxValidatorName].IsValid()
	depPresenter = DepPresenter{
		EmailAddress:     email.String(),
		Domain:           email.Domain(),
		IsFree:           isFree,
		IsSyntax:         isSyntax,
		IsDomain:         validationResults[ev.MXValidatorName].IsValid(),
		IsSmtp:           smtpPresenter.CanConnectSmtp,
		IsVerified:       smtpPresenter.IsDeliverable,
		IsServerDown:     isSyntax && !smtpPresenter.CanConnectSmtp,
		IsGreylisted:     smtpPresenter.IsGreyListed,
		IsDisposable:     !validationResults[ev.DisposableValidatorName].IsValid(),
		IsSuppressed:     !validationResults[ev.BlackListEmailsValidatorName].IsValid(), // TODO find more examples example@example.com
		IsRole:           !validationResults[ev.RoleValidatorName].IsValid(),
		IsHighRisk:       !validationResults[ev.BanWordsUsernameValidatorName].IsValid(), // TODO find more words
		IsCatchall:       smtpPresenter.IsCatchAll,
		TimeTaken:        opts.ExecutedTime(),
		CreditsAvailable: ^uint32(0),
	}

	depPresenter.MailboxvalidatorScore = d.calculateScore(depPresenter)
	depPresenter.Status = depPresenter.MailboxvalidatorScore >= 0.5
	return depPresenter
}

func NewDepValidator() ev.Validator {
	return ev.NewDepBuilder(nil).Set(
		ev.BlackListEmailsValidatorName,
		ev.NewBlackListEmailsValidator(contains.NewSet(hashset.New(
			"example@example.com", "localhost@localhost",
		))),
	).Set(
		ev.BanWordsUsernameValidatorName,
		ev.NewBanWordsUsername(contains.NewInStringsFromArray([]string{"test"})),
	).Set(
		ev.FreeValidatorName,
		ev.FreeDefaultValidator(),
	).Build()
}
