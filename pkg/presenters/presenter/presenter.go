package presenter

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	email "github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
)

type PreparerInterface interface {
	CanProcess(email email.EmailAddressInterface, result ev.ValidationResultInterface) bool
	Process(email email.EmailAddressInterface, result ev.ValidationResultInterface) interface{}
}

type MapPreparers map[ev.ValidatorName]PreparerInterface

func NewMultiplePreparer(processors MapPreparers) MultiplePreparer {
	return MultiplePreparer{processors}
}

type MultiplePreparer struct {
	processors MapPreparers
}

func (p MultiplePreparer) processor(email email.EmailAddressInterface, result ev.ValidationResultInterface) PreparerInterface {
	if processor, ok := p.processors[result.ValidatorName()]; ok && processor.CanProcess(email, result) {
		return processor
	}

	return nil
}

func (p MultiplePreparer) CanProcess(email email.EmailAddressInterface, result ev.ValidationResultInterface) bool {
	return p.processor(email, result) != nil
}

func (p MultiplePreparer) Process(email email.EmailAddressInterface, result ev.ValidationResultInterface) interface{} {
	return p.processor(email, result).Process(email, result)
}
