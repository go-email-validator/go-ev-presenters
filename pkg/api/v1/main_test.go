package v1

import (
	"encoding/json"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evtests"
	openapi "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/go"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator/addition"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/presenter_test"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/prompt_email_verification_api"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/prompt_email_verification_api/cmd/dep_test_generator/struct"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

var valuePresenters = map[preparer.Name]presenter.Interface{
	check_if_email_exist.Name:          nil,
	mailboxvalidator.Name:              nil,
	prompt_email_verification_api.Name: nil,
}

type singleValidationTestArgs struct {
	email      string
	resultType openapi.ResultType
}

type singleValidationTest struct {
	args    singleValidationTestArgs
	want    *openapi.OneOfEmailResponse
	wantErr error
}

func depPresenters(t *testing.T) (tests []singleValidationTest) {
	tests = make([]singleValidationTest, 0)
	rootPath := "../../presenter/"

	{
		fixturePath := rootPath + "/check_if_email_exist/" + presenter_test.DefaultDepFixtureFile
		fixtures := make([]check_if_email_exist.DepPresenter, 0)
		presenter_test.TestDepPresenters(t, &fixtures, fixturePath)
		require.Greater(t, len(fixtures), 0)

		presenters := presenter_test.TestEmailResponses(t, fixturePath, "")
		presenterResult := make(map[string]interface{}, len(fixtures))

		for index, dep := range fixtures {
			tests = append(tests, singleValidationTest{
				args: singleValidationTestArgs{
					email:      dep.Input,
					resultType: openapi.CIEE,
				},
				want: presenters[index],
			})
			presenterResult[dep.Input] = dep
		}
		valuePresenters[check_if_email_exist.Name] = &mockPresenter{presenterResult}
	}
	{
		fixturePath := rootPath + "/mailboxvalidator/" + addition.DepFixtureForViewFile
		fixtures := make([]mailboxvalidator.DepPresenterForView, 0)
		presenter_test.TestDepPresenters(t, &fixtures, fixturePath)
		require.Greater(t, len(fixtures), 0)

		presenters := presenter_test.TestEmailResponses(t, fixturePath, "")
		presenterResult := make(map[string]interface{}, len(fixtures))

		for index, dep := range fixtures {
			tests = append(tests, singleValidationTest{
				args: singleValidationTestArgs{
					email:      dep.EmailAddress,
					resultType: openapi.MAIL_BOX_VALIDATOR,
				},
				want: presenters[index],
			})
			presenterResult[dep.EmailAddress] = dep
		}

		valuePresenters[mailboxvalidator.Name] = &mockPresenter{presenterResult}
	}
	{
		fixturePath := rootPath + "/prompt_email_verification_api/" + presenter_test.DefaultDepFixtureFile
		fixtures := make([]_struct.DepPresenterTest, 0)
		presenter_test.TestDepPresenters(t, &fixtures, fixturePath)
		require.Greater(t, len(fixtures), 0)

		presenters := presenter_test.TestEmailResponses(t, fixturePath, "#.Dep")
		presenterResult := make(map[string]interface{}, len(fixtures))

		for index, dep := range fixtures {
			tests = append(tests, singleValidationTest{
				args: singleValidationTestArgs{
					email:      dep.Email,
					resultType: openapi.PROMPT_EMAIL_VERIFICATION_API,
				},
				want: presenters[index],
			})
			presenterResult[dep.Email] = dep.Dep
		}

		valuePresenters[prompt_email_verification_api.Name] = &mockPresenter{presenterResult}
	}

	return tests
}

func reset() {
	for key := range valuePresenters {
		valuePresenters[key] = nil
	}
}

var opts Options

func TestMain(m *testing.M) {
	getPresenter = func(_ evsmtp.CheckerDTO) presenter.MultiplePresenter {
		return presenter.NewMultiplePresenter(valuePresenters)
	}
	opts = NewOptions()

	var err error
	// Need to correct run of tests
	opts.HTTP.OpenApiPath, err = filepath.Abs("../../../" + opts.HTTP.OpenApiPath)
	if err != nil {
		panic(err)
	}

	server := NewServer(opts)
	err = server.Start()
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()
	evtests.TestMain(m)
}

func TestServer_HTTP(t *testing.T) {
	defer reset()

	tests := depPresenters(t)
	for _, tt := range tests {
		t.Run(tt.args.email+"_"+string(tt.args.resultType), func(t *testing.T) {
			url := "http://" + opts.HTTP.Bind + "/v1/validation/single/" + tt.args.email + "?result_type=" + string(tt.args.resultType)

			client := http.Client{
				Timeout: 10 * time.Second,
			}
			resp, err := client.Get(url)
			require.Equal(t, tt.wantErr, err)
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			require.Nil(t, err)
			got := &openapi.OneOfEmailResponse{}
			err = json.Unmarshal(body, got)
			require.Nil(t, err)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Want\n%v\ngot\n%v", tt.want, got)
			}
		})
	}
}
