package main

import (
	"errors"
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp"
	"github.com/tevino/abool"
	"net/textproto"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var emails []evmail.Address

func getEmail(index int) evmail.Address {
	return emails[index]
}

func main() {
	runtime.GOMAXPROCS(1)

	emails = []evmail.Address{
		//evmail.FromString(mockevmail.ValidEmailString),

		evmail.FromString("none.exist.eamil@gmail.com"),
	}
	depValidator := ev.NewDepBuilder(ev.ValidatorMap{
		ev.SyntaxValidatorName: ev.NewSyntaxValidator(),
		ev.MXValidatorName:     ev.DefaultNewMXValidator(),
		//ev.SMTPValidatorName: ev.GetDefaultSMTPValidator(evsmtp.CheckerDTO{
		//	RandomEmail: func(domain string) (evmail.Address, error) {
		//		return evmail.FromString("w8si10525228lfp@" + domain), nil
		//	},
		//	Options: evsmtp.NewOptions(evsmtp.OptionsDTO{
		//		TimeoutCon:  0,
		//		TimeoutResp: 0,
		//	}),
		//}),
	}).Build()
	//want := ev.NewResult(true, nil, []error{
	//	evsmtp.NewError(evsmtp.RandomRCPTStage, &textproto.Error{
	//		Code: 550,
	//		Msg:  "5.1.1 The email account that you tried to reach does not exist. Please try\n5.1.1 double-checking the recipient's email address for typos or\n5.1.1 unnecessary spaces. Learn more at\n5.1.1  https://support.google.com/mail/?p=NoSuchUser h3si12385421lja.502 - gsmtp",
	//	}),
	//}, ev.SMTPValidatorName)
	// not exists
	want := ev.NewResult(false, []error{errors.New("")}, []error{
		evsmtp.NewError(evsmtp.RandomRCPTStage, &textproto.Error{
			Code: 550,
			Msg:  "5.1.1 The email account that you tried to reach does not exist. Please try\n5.1.1 double-checking the recipient's email address for typos or\n5.1.1 unnecessary spaces. Learn more at\n5.1.1  https://support.google.com/mail/?p=NoSuchUser h3si12385421lja.502 - gsmtp",
		}),
	}, ev.SMTPValidatorName)

	maxIterations := 100000
	wg := sync.WaitGroup{}
	wg.Add(maxIterations)
	start := time.Now()

	needMore := abool.NewBool(true)
	var doneIterations uint64 = 0
	for i := 0; i < maxIterations; i++ {
		if needMore.IsNotSet() {
			wg.Add(-maxIterations)
			break
		}
		go func(i int, email evmail.Address) {
			got := depValidator.Validate(ev.NewInput(email)).(ev.DepValidationResult)
			wg.Done()
			atomic.AddUint64(&doneIterations, 1)
			fmt.Println(doneIterations)

			gotSMTP := got.GetResults()[ev.SMTPValidatorName]

			if gotSMTP.IsValid() != want.IsValid() ||
				gotSMTP.HasErrors() != want.HasErrors() ||
				gotSMTP.HasWarnings() != want.HasWarnings() {
				needMore.UnSet()
				fmt.Errorf("iteration = %v, Errors() = %v, want %v", i, gotSMTP, want)
			}
		}(i, getEmail(i%len(emails)))
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}
