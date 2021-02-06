package common

import (
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MaxRequestPerOneEmail = 20000
	MaxInt                = 2147483647
)

func getEmail(emails []evmail.Address, index int) evmail.Address {
	return emails[index%len(emails)]
}

func EmailsTest(emails []evmail.Address, maxIterations int, want ev.ValidationResult) {
	if maxIterations == 0 {
		maxIterations = len(emails)
	}

	depValidator := ev.NewDepBuilder(ev.ValidatorMap{
		ev.SyntaxValidatorName: ev.NewSyntaxValidator(),
		ev.MXValidatorName:     ev.DefaultNewMXValidator(),
		/*ev.SMTPValidatorName: ev.GetDefaultSMTPValidator(evsmtp.CheckerDTO{
			RandomEmail: func(domain string) (evmail.Address, error) {
				return evmail.FromString("w8si10525228lfp@" + domain), nil
			},
			Options: evsmtp.NewOptions(evsmtp.OptionsDTO{
				TimeoutCon:  20 * time.Second,
				TimeoutResp: 20 * time.Second,
			}),
		}),*/
	}).Build()

	wg := sync.WaitGroup{}
	wg.Add(maxIterations)
	start := time.Now()

	done := make(chan struct{}, 1)
	var doneIterations uint64 = 0
	for i := 0; i < maxIterations; i++ {
		go func(i int, email evmail.Address) {
			defer wg.Done()
			got := depValidator.Validate(ev.NewInput(email)).(ev.DepValidationResult)
			atomic.AddUint64(&doneIterations, 1)
			fmt.Println(doneIterations)

			gotSMTP := got.GetResults()[ev.SMTPValidatorName]

			if gotSMTP.IsValid() != want.IsValid() ||
				gotSMTP.HasErrors() != want.HasErrors() ||
				gotSMTP.HasWarnings() != want.HasWarnings() {
				defer close(done)
				fmt.Println(email.String(), doneIterations)
				fmt.Println(fmt.Errorf("iteration = %v, Errors() = %v, want %v", i, gotSMTP, want))
			}
		}(i, getEmail(emails, i))
	}

	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-done:
	}

	fmt.Println(time.Since(start))
}
