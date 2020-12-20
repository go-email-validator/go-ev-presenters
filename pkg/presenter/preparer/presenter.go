package preparer

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"time"
)

type Name string

type OptionsInterface interface {
	IsOptions()
}

type TimeOptions interface {
	OptionsInterface
	ExecutedTime() time.Duration
}

type Options struct {
	ExecutedTimeValue time.Duration
}

func (_ Options) IsOptions() {}
func (o Options) ExecutedTime() time.Duration {
	return o.ExecutedTimeValue
}

type Interface interface {
	CanPrepare(email email.EmailAddressInterface, result ev.ValidationResultInterface, opts OptionsInterface) bool
	Prepare(email email.EmailAddressInterface, result ev.ValidationResultInterface, opts OptionsInterface) interface{}
}

type MapPreparers map[ev.ValidatorName]Interface

func NewMultiplePreparer(preparers MapPreparers) MultiplePreparer {
	return MultiplePreparer{preparers}
}

type MultiplePreparer struct {
	preparers MapPreparers
}

func (p MultiplePreparer) preparer(email email.EmailAddressInterface, result ev.ValidationResultInterface, opts OptionsInterface) Interface {
	if preparer, ok := p.preparers[result.ValidatorName()]; ok && preparer.CanPrepare(email, result, opts) {
		return preparer
	}

	return nil
}

func (p MultiplePreparer) CanPrepare(email email.EmailAddressInterface, result ev.ValidationResultInterface, opts OptionsInterface) bool {
	return p.preparer(email, result, opts) != nil
}

func (p MultiplePreparer) Prepare(email email.EmailAddressInterface, result ev.ValidationResultInterface, opts OptionsInterface) interface{} {
	return p.preparer(email, result, opts).Prepare(email, result, opts)
}
