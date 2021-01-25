package v1

import (
	"encoding/json"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp"
	mockev "github.com/go-email-validator/go-email-validator/test/mock/ev"
	mockevmail "github.com/go-email-validator/go-email-validator/test/mock/ev/evmail"
	openapi "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/go"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"runtime"
	"testing"
	"time"
)

func newMultiplePresentersDefault(b *testing.B, checkerDTO evsmtp.CheckerDTO, _ Options) presenter.MultiplePresenter {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	checker := evsmtp.NewChecker(checkerDTO)
	ev.NewSMTPValidator(checker)

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

	mxValidator := ev.DefaultNewMXValidator()
	mxMock := mockev.NewMockValidator(ctrl)
	mxMock.EXPECT().GetDeps().Return(make([]ev.ValidatorName, 0)).AnyTimes()
	mxMock.EXPECT().Validate(gomock.Any()).
		DoAndReturn(func(input ev.Input, _ ...ev.ValidationResult) ev.ValidationResult {
			mxValidator.Validate(input)

			return ev.NewMXValidationResult(evsmtp.MXs{
				{
					Host: "127.0.0.1",
					Pref: 0,
				},
			}, ev.NewValidResult(ev.MXValidatorName).(*ev.AValidationResult))
		}).
		AnyTimes()

	builder := ev.NewDepBuilder(nil)

	validator := builder.Set(ev.MXValidatorName, mxMock).Set(
		ev.SyntaxValidatorName,
		ev.NewSyntaxRegexValidator(nil),
	).Set(
		ev.SMTPValidatorName,
		smtpValidator,
	).Build()

	return presenter.NewMultiplePresenter(map[preparer.Name]presenter.Interface{
		check_if_email_exist.Name: presenter.NewPresenter(
			evmail.FromString,
			validator,
			check_if_email_exist.NewDepPreparerDefault(),
		),
	})
}

func BenchmarkServer(b *testing.B) {
	runtime.GOMAXPROCS(1) // Use only wan processor

	getPresenter = func(checkerDTO evsmtp.CheckerDTO, opts Options) presenter.MultiplePresenter {
		return newMultiplePresentersDefault(b, checkerDTO, opts)
	}

	opts := NewOptions()
	opts.IsVerbose = true
	_, opts = startServer(&opts)

	<-time.After(1 * time.Millisecond)

	email := mockevmail.GetValidTestEmail()
	reqURL := "http://" + opts.HTTP.Bind + "/v1/validation/single/" + email.String()
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	b.ResetTimer()
	b.Run("Parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				resp, err := client.Get(reqURL)
				require.Nil(b, err)
				if err != nil {
					a := 1
					_ = a
				}
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				require.Nil(b, err)
				got := &openapi.OneOfEmailResponse{}
				err = json.Unmarshal(body, got)
				require.Nil(b, err)
			}
		})
	})
}
