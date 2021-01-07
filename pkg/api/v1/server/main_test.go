package server

import (
	"context"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evtests"
	v1 "github.com/go-email-validator/go-ev-presenters/pkg/api/v1"
	apiciee "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/check_if_email_exist"
	apimbv "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/mailboxvalidator"
	api_prompt_email_verification "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/prompt_email_verification_api"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator/addition"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/prompt_email_verification_api"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/prompt_email_verification_api/cmd/dep_test_generator/struct"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	resultType v1.ResultType
}

type singleValidationTest struct {
	args    singleValidationTestArgs
	want    *v1.EmailResponse
	wantErr error
}

func depPresenters(t *testing.T) (tests []singleValidationTest) {
	tests = make([]singleValidationTest, 0)
	rootPath := "../../../presenter/"

	{
		fixturePath := rootPath + "/check_if_email_exist/" + common.DefaultDepFixtureFile
		fixtures := make([]check_if_email_exist.DepPresenter, 0)
		common.TestDepPresenters(t, &fixtures, fixturePath)
		require.Greater(t, len(fixtures), 0)

		presenters := common.TestEmailResponses(t, &apiciee.Result{}, fixturePath, "")
		presenterResult := make(map[string]interface{}, len(fixtures))

		for index, dep := range fixtures {
			tests = append(tests, singleValidationTest{
				args: singleValidationTestArgs{
					email:      dep.Input,
					resultType: v1.ResultType_CIEE,
				},
				want: &v1.EmailResponse{Result: &v1.EmailResponse_CheckIfEmailExist{
					CheckIfEmailExist: presenters[index].(*apiciee.Result),
				}},
			})
			presenterResult[dep.Input] = dep
		}
		valuePresenters[check_if_email_exist.Name] = &mockPresenter{presenterResult}
	}
	{
		fixturePath := rootPath + "/mailboxvalidator/" + addition.DepFixtureForViewFile
		fixtures := make([]mailboxvalidator.DepPresenterForView, 0)
		common.TestDepPresenters(t, &fixtures, fixturePath)
		require.Greater(t, len(fixtures), 0)

		presenters := common.TestEmailResponses(t, &apimbv.Result{}, fixturePath, "")
		presenterResult := make(map[string]interface{}, len(fixtures))

		for index, dep := range fixtures {
			tests = append(tests, singleValidationTest{
				args: singleValidationTestArgs{
					email:      dep.EmailAddress,
					resultType: v1.ResultType_MAIL_BOX_VALIDATOR,
				},
				want: &v1.EmailResponse{Result: &v1.EmailResponse_MailBoxValidator{
					MailBoxValidator: presenters[index].(*apimbv.Result),
				}},
			})
			presenterResult[dep.EmailAddress] = dep
		}

		valuePresenters[mailboxvalidator.Name] = &mockPresenter{presenterResult}
	}
	{
		fixturePath := rootPath + "/prompt_email_verification_api/" + common.DefaultDepFixtureFile
		fixtures := make([]_struct.DepPresenterTest, 0)
		common.TestDepPresenters(t, &fixtures, fixturePath)
		require.Greater(t, len(fixtures), 0)

		presenters := common.TestEmailResponses(t, &api_prompt_email_verification.Result{}, fixturePath, "#.Dep")
		presenterResult := make(map[string]interface{}, len(fixtures))

		for index, dep := range fixtures {
			tests = append(tests, singleValidationTest{
				args: singleValidationTestArgs{
					email:      dep.Email,
					resultType: v1.ResultType_PROMPT_EMAIL_VERIFICATION_API,
				},
				want: &v1.EmailResponse{Result: &v1.EmailResponse_PromptEmailVerificationApi{
					PromptEmailVerificationApi: presenters[index].(*api_prompt_email_verification.Result),
				}},
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
	getPresenter = func() presenter.MultiplePresenter {
		return presenter.NewMultiplePresenter(valuePresenters)
	}
	opts = NewOptions()

	server := NewServer(opts)
	err := server.Start()
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()
	evtests.TestMain(m)
}

func TestServer_GRPC(t *testing.T) {
	defer reset()

	conn, err := grpc.Dial(opts.GRPC.Bind, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Printf("did not connect: %v", err)
		require.True(t, false)
	}
	defer conn.Close()
	c := v1.NewEmailValidationClient(conn)

	tests := depPresenters(t)

	for _, tt := range tests {
		t.Run(tt.args.email+"_"+strconv.Itoa(int(tt.args.resultType)), func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			got, err := c.SingleValidation(ctx, &v1.EmailRequest{
				Email:      tt.args.email,
				ResultType: tt.args.resultType,
			})

			require.Nil(t, err)
			if !proto.Equal(proto.Message(tt.want), got) {
				t.Errorf("Want\n%v\ngot\n%v", tt.want.GetResult(), got.GetResult())
			}
		})
	}
}

func TestServer_HTTP(t *testing.T) {
	defer reset()

	tests := depPresenters(t)
	for _, tt := range tests {
		t.Run(tt.args.email+"_"+strconv.Itoa(int(tt.args.resultType)), func(t *testing.T) {
			url := "http://" + opts.HTTP.Bind + "/v1/validation/single/" + tt.args.email + "?result_type=" + strconv.Itoa(int(tt.args.resultType))

			client := http.Client{
				Timeout: 10 * time.Second,
			}
			resp, err := client.Get(url)
			require.Equal(t, tt.wantErr, err)
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			require.Nil(t, err)
			got := &v1.EmailResponse{}
			err = protojson.Unmarshal(body, got)
			require.Nil(t, err)
			if !proto.Equal(proto.Message(tt.want), got) {
				t.Errorf("Want\n%v\ngot\n%v", tt.want.GetResult(), got.GetResult())
			}
		})
	}
}
