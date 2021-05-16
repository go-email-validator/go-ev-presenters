package v1

import (
	"encoding/json"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evtests"
	openapi "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/go"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/converter"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/mailboxvalidator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/mailboxvalidator/addition"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/presentation_test"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/prompt_email_verification_api"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/prompt_email_verification_api/cmd/dep_test_generator/struct"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"
)

var valuePresenters = map[converter.Name]presentation.Interface{
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
	want    *openapi.EmailResponse
	wantErr error
}

func depPresenters(t *testing.T) (tests []singleValidationTest) {
	tests = make([]singleValidationTest, 0)
	rootPath := "../../presentation/"

	{
		fixturePath := rootPath + "/check_if_email_exist/" + presentation_test.DefaultDepFixtureFile
		fixtures := make([]check_if_email_exist.DepPresentation, 0)
		presentation_test.TestDepPresentations(t, &fixtures, fixturePath)
		require.Greater(t, len(fixtures), 0)

		presenters := presentation_test.TestEmailResponses(t, fixturePath, "", func(data []byte) *openapi.EmailResponse {
			result := openapi.CheckIfEmailExistResult{}
			err := json.Unmarshal(data, &result)
			require.Nil(t, err, fixturePath, string(data))

			return &openapi.EmailResponse{
				CheckIfEmailExist: result,
			}
		})
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
		fixtures := make([]mailboxvalidator.DepPresentationForView, 0)
		presentation_test.TestDepPresentations(t, &fixtures, fixturePath)
		require.Greater(t, len(fixtures), 0)

		presenters := presentation_test.TestEmailResponses(t, fixturePath, "", func(data []byte) *openapi.EmailResponse {
			result := openapi.MailboxvalidatorResult{}
			err := json.Unmarshal(data, &result)
			require.Nil(t, err, fixturePath, string(data))

			return &openapi.EmailResponse{
				Mailboxvalidator: result,
			}
		})
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
		fixturePath := rootPath + "/prompt_email_verification_api/" + presentation_test.DefaultDepFixtureFile
		fixtures := make([]_struct.DepPresentationTest, 0)
		presentation_test.TestDepPresentations(t, &fixtures, fixturePath)
		require.Greater(t, len(fixtures), 0)

		presenters := presentation_test.TestEmailResponses(t, fixturePath, "#.Dep", func(data []byte) *openapi.EmailResponse {
			result := openapi.PromptEmailVerificationApiResult{}
			err := json.Unmarshal(data, &result)
			require.Nil(t, err, fixturePath, string(data))

			return &openapi.EmailResponse{
				PromptEmailVerificationApi: result,
			}
		})
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

func startServer(InOpts *Options) (*Server, Options) {
	opts := NewOptions()
	if InOpts != nil {
		opts = *InOpts
	}

	opts.Fiber.IdleTimeout = 500 * time.Millisecond
	opts.HTTP.OpenApiResponseValidation = true

	server := NewServer(DefaultFiberFactory, opts)
	err := server.Start()
	if err != nil {
		panic(err)
	}

	return &server, opts
}

func shutdownServer(server *Server) {
	server.Shutdown()
}

func TestMain(m *testing.M) {
	evtests.TestMain(m)
}

func TestServer_HTTP(t *testing.T) {
	getPresenter = func(_ evsmtp.CheckerDTO, _ Options) presentation.MultiplePresenter {
		return presentation.NewMultiplePresenter(valuePresenters)
	}
	server, opts := startServer(nil)
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
				require.False(t, strings.Contains(string(body), "{\"message\":\""), "Invalid answer \n"+string(body))

				got := &openapi.EmailResponse{}
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

	server, opts := startServer(nil)
	defer shutdownServer(server)

	// Some data or functional cannot be matched, see more nearby DepPresentation of emails
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
					Timeout: 15 * time.Second,
				}
				resp, err := client.Get(url)
				require.Equal(t, tt.wantErr, err)
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				require.Nil(t, err)
				require.False(t, strings.Contains(string(body), "{\"message\":\""))
				got := &openapi.EmailResponse{}
				err = json.Unmarshal(body, got)
				require.Nil(t, err)

				sort.Strings(got.CheckIfEmailExist.Mx.Records)
				sort.Strings(got.PromptEmailVerificationApi.MxRecords.Records)
				got.Mailboxvalidator.TimeTaken = "0"

				sort.Strings(tt.want.CheckIfEmailExist.Mx.Records)
				sort.Strings(tt.want.PromptEmailVerificationApi.MxRecords.Records)
				tt.want.Mailboxvalidator.TimeTaken = "0"

				// TODO create more complex approach to skip some problem on outside service
				if name == "tvzamhkdc@emlhub.com_PROMPT_EMAIL_VERIFICATION_API" {
					got.PromptEmailVerificationApi.CanConnectSmtp = false
				}

				if !reflect.DeepEqual(tt.want, got) {
					t.Errorf("Want\n%v\ngot\n%v", tt.want, got)
				}
			})
		}
	})
}
