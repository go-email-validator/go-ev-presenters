package cmd

import (
	v1 "github.com/go-email-validator/go-ev-presenters/pkg/api/v1"
	"os"
)

const (
	httpBindFlag            = "http-bind"
	verboseFlag             = "verbose"
	headersFlag             = "headers"
	ipsFlag                 = "ips"
	smtpProxyFlag           = "smtp-proxy"
	helloNameFlag           = "localname"
	memCachedFlag           = "memcached"
	ristrettoFlag           = "ristretto"
	fiberStartupMessageFlag = "fiber-startup-msg"
	showOpenApiFlag         = "openapi"
)

var fiberStartupMessage bool

func init() {
	// For Heroku
	if gatewayPort := os.Getenv("PORT"); gatewayPort != "" {
		opts.HTTP.Bind = "0.0.0.0:" + gatewayPort
	}

	rootCmd.Flags().StringVar(&opts.HTTP.Bind, httpBindFlag, opts.HTTP.Bind, "HTTP bind address")

	_, opts.IsVerbose = os.LookupEnv(v1.VerboseEnv)
	rootCmd.Flags().BoolVarP(&opts.IsVerbose, verboseFlag, "v", opts.IsVerbose, "Show DEBUG log information")

	rootCmd.Flags().StringToStringVar(&opts.Auth.Headers, headersFlag, opts.Auth.Headers, "Map headers values")

	rootCmd.Flags().StringArrayVar(&opts.Auth.IPs, ipsFlag, opts.Auth.IPs, "List acceptable ips")

	rootCmd.Flags().StringVar(&opts.Validator.SMTPProxy, smtpProxyFlag, opts.Validator.SMTPProxy, "Proxy for smtp calling")

	if helloName := os.Getenv(v1.HelloNameEnv); helloName != "" {
		opts.Validator.HelloName = helloName
	}

	rootCmd.Flags().StringArrayVar(&opts.Validator.Memcached, memCachedFlag, opts.Validator.Memcached, "List of memcached servers")

	rootCmd.Flags().BoolVar(&opts.Validator.Ristretto, ristrettoFlag, opts.Validator.Ristretto, "Is Ristretto active?")

	rootCmd.Flags().StringVar(&opts.Validator.HelloName, helloNameFlag, opts.Validator.HelloName, "HelloName for SMTP HELO command")

	rootCmd.Flags().BoolVar(&fiberStartupMessage, fiberStartupMessageFlag, !opts.Fiber.DisableStartupMessage, "Show or not Fiber startup message")

	rootCmd.Flags().BoolVar(&opts.HTTP.ShowOpenApi, showOpenApiFlag, opts.HTTP.ShowOpenApi, "Show OpenApi")

	opts.Fiber.DisableStartupMessage = !fiberStartupMessage
}
