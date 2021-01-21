package v1

import (
	"encoding/json"
	"github.com/emirpasic/gods/sets/hashset"
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
	"sort"
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

func startServer() *Server {
	opts = NewOptions()

	var err error
	// Need to correct run of tests
	opts.HTTP.OpenApiPath, err = filepath.Abs("../../../" + opts.HTTP.OpenApiPath)
	if err != nil {
		panic(err)
	}

	opts.Fiber.IdleTimeout = 500 * time.Millisecond

	server := NewServer(opts)
	err = server.Start()
	if err != nil {
		panic(err)
	}

	return &server
}

func shutdownServer(server *Server) {
	server.Shutdown()
}

func TestMain(m *testing.M) {
	evtests.TestMain(m)
}

func TestServer_HTTP(t *testing.T) {
	getPresenter = func(_ evsmtp.CheckerDTO) presenter.MultiplePresenter {
		return presenter.NewMultiplePresenter(valuePresenters)
	}
	server := startServer()
	defer reset()
	defer shutdownServer(server)

	tests := depPresenters(t)
	t.Run("Parallel", func(t *testing.T) {
		for _, tt := range tests {
			tt := tt
			t.Run(tt.args.email+"_"+string(tt.args.resultType), func(t *testing.T) {
				t.Parallel()
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
	})
}

func TestServer_HTTP_FUNC(t *testing.T) {
	evtests.FunctionalSkip(t)

	server := startServer()
	defer shutdownServer(server)

	// Some data or functional cannot be matched, see more nearby DepPresenter of emails
	skipEmail := hashset.New(
		// TODO problem with SMTP, CIEE think that email is not is_catch_all. Need to run and research source code on RUST
		"sewag33689@itymail.com",
		/* TODO add proxy to test
		5.7.1 Service unavailable, Client host [94.181.152.110] blocked using Spamhaus. To request removal from this list see https://www.spamhaus.org/query/ip/94.181.152.110 (AS3130). [BN8NAM12FT053.eop-nam12.prod.protection.outlook.com]
		*/
		"salestrade86@hotmail.com",
		"monicaramirezrestrepo@hotmail.com",
		// TODO CIEE banned
		"credit@mail.ru",
		// TODO need check source code
		"asdasd@tradepro.net",
		// TODO need check source code
		"y-numata@senko.ed.jp",
	)

	skipNames := hashset.New(
		"zxczxczxc@joycasinoru_MAIL_BOX_VALIDATOR",
		"derduzikne@nedoz.com_MAIL_BOX_VALIDATOR",
		"tvzamhkdc@emlhub.com_MAIL_BOX_VALIDATOR",
		"admin@gmail.com_MAIL_BOX_VALIDATOR",
		"pr@yandex-team.ru_MAIL_BOX_VALIDATOR",
		"tvzamhkdc@emlhub.com_PROMPT_EMAIL_VERIFICATION_API",
	)

	tests := depPresenters(t)
	t.Run("Parallel", func(t *testing.T) {
		for _, tt := range tests {
			tt := tt
			name := tt.args.email + "_" + string(tt.args.resultType)
			if skipEmail.Contains(tt.args.email) || skipNames.Contains(name) {
				t.Logf("skipped %v", name)
				continue
			}
			t.Run(name, func(t *testing.T) {
				t.Parallel()
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

				sort.Strings(got.MxRecords.Records)
				sort.Strings(got.Mx.Records)
				got.TimeTaken = "0"

				sort.Strings(tt.want.MxRecords.Records)
				sort.Strings(tt.want.Mx.Records)
				tt.want.TimeTaken = "0"

				// TODO create more complex approach to skip some problem on outside service
				if name == "tvzamhkdc@emlhub.com_PROMPT_EMAIL_VERIFICATION_API" {
					got.PromptEmailVerificationApiResult.CanConnectSmtp = false
				}

				if !reflect.DeepEqual(tt.want, got) {
					t.Errorf("Want\n%v\ngot\n%v", tt.want, got)
				}
			})
		}
	})
}
