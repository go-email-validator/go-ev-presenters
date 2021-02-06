package v1

import (
	"encoding/csv"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp"
	"github.com/go-email-validator/go-email-validator/pkg/ev/utils"
	openapi "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/go"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/prompt_email_verification_api"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	HTTPDefaultHost = "0.0.0.0:8090"
	SwaggerPath     = "/api/v1/openapiv3/ev.openapiv3.yaml"

	EnvPrefix              = "EV_"
	HttpBindEnv            = EnvPrefix + "HTTP_BIND"
	VerboseEnv             = EnvPrefix + "VERBOSE"
	HeadersEnv             = EnvPrefix + "HEADERS"
	IPsEnv                 = EnvPrefix + "IPS"
	SMTPProxyEnv           = EnvPrefix + "SMTP_PROXY"
	MemcachedEnv           = EnvPrefix + "MEMCACHED"
	RistrettoEnv           = EnvPrefix + "RISTRETTO"
	HelloNameEnv           = EnvPrefix + "HELLONAME"
	FiberStartupMessageEnv = EnvPrefix + "FIBER_STARTUP_MSG"
	ShowOpenApiEnv         = EnvPrefix + "OPENAPI"
)

func DefaultInstance(opts Options) openapi.EmailValidationApiRouter {
	authenticationFunc := AuthFuncComplex(opts.Auth)
	if authenticationFunc == nil {
		authenticationFunc = openapi3filter.NoopAuthenticationFunc
	}

	defaultCheckerOptions := evsmtp.DefaultOptions()

	return NewEmailValidationApiController(EmailValidationApiControllerDTO{
		Presenter: getPresenter(evsmtp.CheckerDTO{
			SendMailFactory: evsmtp.NewSendMailFactory(evsmtp.H12IODial, nil),
			Options: evsmtp.NewOptions(evsmtp.OptionsDTO{
				EmailFrom:   defaultCheckerOptions.EmailFrom(),
				HelloName:   utils.DefaultString(opts.Validator.HelloName, defaultCheckerOptions.HelloName()),
				Proxy:       utils.DefaultString(opts.Validator.SMTPProxy, defaultCheckerOptions.Proxy()),
				TimeoutCon:  defaultCheckerOptions.TimeoutConnection(),
				TimeoutResp: defaultCheckerOptions.TimeoutResponse(),
				Port:        defaultCheckerOptions.Port(),
			}),
		}, opts),
		Matching: map[openapi.ResultType]preparer.Name{
			openapi.CIEE:                          check_if_email_exist.Name,
			openapi.CHECK_IF_EMAIL_EXIST:          check_if_email_exist.Name,
			openapi.MBV:                           mailboxvalidator.Name,
			openapi.MAILBOXVALIDATOR:              mailboxvalidator.Name,
			openapi.MAIL_BOX_VALIDATOR:            mailboxvalidator.Name,
			openapi.PEVA:                          prompt_email_verification_api.Name,
			openapi.PROMPT_EMAIL_VERIFICATION_API: prompt_email_verification_api.Name,
		},
		OpenApiValidator: openapi.NewValidator(
			openapi.RouterFromPath(opts.HTTP.OpenApiPath),
			&openapi3filter.Options{
				AuthenticationFunc: authenticationFunc,
			},
		),
	})
}

var getPresenter = NewMultiplePresentersDefault

func NewOptions() Options {
	return Options{
		IsVerbose: false,
		Validator: NewValidator(),
		HTTP:      NewHTTPOptions(),
		Auth:      NewAuthOptions(),
		Fiber: fiber.Config{
			DisableStartupMessage: true,
			// TODO move to cli config
			IdleTimeout:  5 * time.Second,
			WriteTimeout: 20 * time.Second,
			ReadTimeout:  20 * time.Second,
		},
	}
}

func OptionsFromEnvironment() Options {
	opts := NewOptions()

	if bind := os.Getenv(HttpBindEnv); bind != "" {
		opts.HTTP.Bind = bind
	}

	if headers := os.Getenv(HeadersEnv); headers != "" {
		opts.Auth.Headers = stringToStringConvSilence(headers)
	}

	if ips := os.Getenv(IPsEnv); ips != "" {
		opts.Auth.IPs = readAsCSVSilence(ips)
	}

	if smtpProxy := os.Getenv(SMTPProxyEnv); smtpProxy != "" {
		opts.Validator.SMTPProxy = smtpProxy
	}

	if memCached := os.Getenv(MemcachedEnv); memCached != "" {
		opts.Validator.Memcached = readAsCSVSilence(memCached)
	}

	if ristretto, hasRistretto := os.LookupEnv(RistrettoEnv); hasRistretto {
		if ristrettoBool, err := strconv.ParseBool(ristretto); err == nil {
			opts.Validator.Ristretto = ristrettoBool
		}
	}

	if fiberStartupMessage, hasFiberStartupMessage := os.LookupEnv(FiberStartupMessageEnv); hasFiberStartupMessage {
		if fiberStartupMessageBool, err := strconv.ParseBool(fiberStartupMessage); err == nil {
			opts.Fiber.DisableStartupMessage = !fiberStartupMessageBool
		}
	}

	if showOpenApi, hasShowOpenApi := os.LookupEnv(ShowOpenApiEnv); hasShowOpenApi {
		if showOpenApiBool, err := strconv.ParseBool(showOpenApi); err == nil {
			opts.HTTP.ShowOpenApi = showOpenApiBool
		}
	}

	return opts
}

type Options struct {
	IsVerbose bool
	Validator Validator
	HTTP      HTTPOptions
	Auth      AuthOptions
	Fiber     fiber.Config
}

func NewValidator() Validator {
	return Validator{}
}

type Validator struct {
	SMTPProxy string
	HelloName string
	Memcached []string
	Ristretto bool
}

func NewHTTPOptions() HTTPOptions {
	return HTTPOptions{
		Bind:        HTTPDefaultHost,
		OpenApiPath: SwaggerPath,
	}
}

type HTTPOptions struct {
	Bind        string
	ShowOpenApi bool
	OpenApiPath string
}

func NewAuthOptions() AuthOptions {
	return AuthOptions{}
}

type AuthOptions struct {
	Headers map[string]string
	IPs     []string
}

func readAsCSV(val string) ([]string, error) {
	if val == "" {
		return []string{}, nil
	}
	stringReader := strings.NewReader(val)
	csvReader := csv.NewReader(stringReader)
	return csvReader.Read()
}

func readAsCSVSilence(val string) []string {
	result, errs := readAsCSV(val)

	if errs != nil {
		panic(errs)
	}

	return result
}

func stringToStringConv(val string) (map[string]string, error) {
	val = strings.Trim(val, "[]")
	// An empty string would cause an empty map
	if len(val) == 0 {
		return map[string]string{}, nil
	}
	r := csv.NewReader(strings.NewReader(val))
	ss, err := r.Read()
	if err != nil {
		return nil, err
	}
	out := make(map[string]string, len(ss))
	for _, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("%s must be formatted as key=value", pair)
		}
		out[kv[0]] = kv[1]
	}
	return out, nil
}

func stringToStringConvSilence(val string) map[string]string {
	result, errs := stringToStringConv(val)

	if errs != nil {
		panic(errs)
	}

	return result
}
