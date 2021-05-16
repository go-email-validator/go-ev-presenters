package v1

import (
	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/store"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/converter"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/mailboxvalidator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/prompt_email_verification_api"
	"time"
)

func NewMultiplePresentersDefault(checkerDTO evsmtp.CheckerDTO, opts Options) presentation.ValidationPresenter {
	var cache CacheInterface

	withMemcached := len(opts.Validator.Memcached) > 0
	withRistretto := opts.Validator.Ristretto
	hasCache := withMemcached || withRistretto

	cacheCheckerOpts := &store.Options{
		Expiration: 10 * time.Minute,
	}
	cacheValidatorOpts := &store.Options{
		Expiration: 30 * time.Minute,
	}
	if withMemcached {
		cache = NewMemcachedCache(opts.Validator.Memcached, cacheCheckerOpts, cacheValidatorOpts)
	} else if withRistretto {
		var checkerCount int64 = 10000
		var validatorCount int64 = 50000

		cache = NewRistrettoCache(
			&ristretto.Config{
				NumCounters: 10 * checkerCount,
				MaxCost:     400 * checkerCount,
				BufferItems: 64,
			},
			&ristretto.Config{
				NumCounters: 10 * validatorCount,
				MaxCost:     500 * validatorCount,
				BufferItems: 64,
			}, cacheCheckerOpts, cacheValidatorOpts)
	}

	checker := evsmtp.NewChecker(checkerDTO)
	if hasCache {
		checker = cache.Checker(checker)
	}

	smtpValidator := ev.NewWarningsDecorator(
		ev.NewSMTPValidator(checker),
		ev.NewIsWarning(hashset.New(evsmtp.RandomRCPTStage), func(warningMap ev.WarningSet) ev.IsWarning {
			return func(err error) bool {
				errSMTP, ok := err.(evsmtp.Error)
				if !ok {
					return false
				}
				return warningMap.Contains(errSMTP.Stage())
			}
		}),
	)
	if hasCache {
		smtpValidator = cache.Validator(smtpValidator)
	}

	return presentation.NewValidationPresenter(map[converter.Name]presentation.Interface{
		check_if_email_exist.Name: presentation.NewPresenter(
			evmail.FromString,
			check_if_email_exist.NewDepValidator(smtpValidator),
			check_if_email_exist.NewDepConverterDefault(),
		),
		mailboxvalidator.Name: presentation.NewPresenter(
			mailboxvalidator.EmailFromString,
			mailboxvalidator.NewDepValidator(smtpValidator),
			mailboxvalidator.NewDepConverterForViewDefault(),
		),
		prompt_email_verification_api.Name: presentation.NewPresenter(
			prompt_email_verification_api.EmailFromString,
			prompt_email_verification_api.NewDepValidator(smtpValidator),
			prompt_email_verification_api.NewDepConverterDefault(),
		),
	})
}
