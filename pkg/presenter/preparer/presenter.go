package preparer

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
)

type Name string

type Interface interface {
	CanPrepare(email email.EmailAddressInterface, result ev.ValidationResultInterface) bool
	Prepare(email email.EmailAddressInterface, result ev.ValidationResultInterface) interface{}
}

type MapPreparers map[ev.ValidatorName]Interface

func NewMultiplePreparer(preparers MapPreparers) MultiplePreparer {
	return MultiplePreparer{preparers}
}

type MultiplePreparer struct {
	preparers MapPreparers
}

func (p MultiplePreparer) preparer(email email.EmailAddressInterface, result ev.ValidationResultInterface) Interface {
	if preparer, ok := p.preparers[result.ValidatorName()]; ok && preparer.CanPrepare(email, result) {
		return preparer
	}

	return nil
}

func (p MultiplePreparer) CanPrepare(email email.EmailAddressInterface, result ev.ValidationResultInterface) bool {
	return p.preparer(email, result) != nil
}

func (p MultiplePreparer) Prepare(email email.EmailAddressInterface, result ev.ValidationResultInterface) interface{} {
	return p.preparer(email, result).Prepare(email, result)
}
