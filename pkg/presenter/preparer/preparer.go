package preparer

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"time"
)

type Name string

type Options interface {
	IsOptions()
	ExecutedTime() time.Duration
}

func NewOptions(executedTime time.Duration) Options {
	return options{
		ExecutedTimeValue: executedTime,
	}
}

type options struct {
	ExecutedTimeValue time.Duration
}

func (_ options) IsOptions() {}
func (o options) ExecutedTime() time.Duration {
	return o.ExecutedTimeValue
}

type Interface interface {
	CanPrepare(email email.EmailAddress, result ev.ValidationResult, opts Options) bool
	Prepare(email email.EmailAddress, result ev.ValidationResult, opts Options) interface{}
}

type MapPreparers map[ev.ValidatorName]Interface

func NewMultiplePreparer(preparers MapPreparers) MultiplePreparer {
	return MultiplePreparer{preparers}
}

type MultiplePreparer struct {
	preparers MapPreparers
}

func (p MultiplePreparer) preparer(email email.EmailAddress, result ev.ValidationResult, opts Options) Interface {
	if preparer, ok := p.preparers[result.ValidatorName()]; ok && preparer.CanPrepare(email, result, opts) {
		return preparer
	}

	return nil
}

func (p MultiplePreparer) CanPrepare(email email.EmailAddress, result ev.ValidationResult, opts Options) bool {
	return p.preparer(email, result, opts) != nil
}

func (p MultiplePreparer) Prepare(email email.EmailAddress, result ev.ValidationResult, opts Options) interface{} {
	return p.preparer(email, result, opts).Prepare(email, result, opts)
}
