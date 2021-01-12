package cmd

import (
	"os"
)

const (
	envPrefix               = "EV_"
	httpBindFlag            = "http-bind"
	httpBindEnv             = envPrefix + "HTTP_BIND"
	verboseFlag             = "verbose"
	verboseEnv              = envPrefix + "VERBOSE"
	apiKeyFlag              = "api-key"
	apiKeyEnv               = envPrefix + "API_KEY"
	smtpProxyFlag           = "smtp-proxy"
	smtpProxyEnv            = envPrefix + "SMTP_PROXY"
	localNameFlag           = "localname"
	localNameEnv            = envPrefix + "LOCALNAME"
	fiberStartupMessageFlag = "fiber-startup-msg"
	fiberStartupMessageEnv  = envPrefix + "FIBER_STARTUP_MSG"
)

var isVerbose bool
var fiberStartupMessage bool

func init() {
	// For Heroku
	if gatewayPort := os.Getenv("PORT"); gatewayPort != "" {
		opts.HTTP.Bind = "0.0.0.0:" + gatewayPort
	}

	if bind := os.Getenv(httpBindEnv); bind != "" {
		opts.HTTP.Bind = bind
	}
	rootCmd.Flags().StringVar(&opts.HTTP.Bind, httpBindFlag, opts.HTTP.Bind, "HTTP bind address")

	_, isVerbose = os.LookupEnv(verboseEnv)
	rootCmd.Flags().BoolVarP(&isVerbose, verboseFlag, "v", isVerbose, "Show DEBUG log information")

	if apiKey := os.Getenv(apiKeyEnv); apiKey != "" {
		opts.Auth.Key = apiKey
	}
	rootCmd.Flags().StringVarP(&opts.Auth.Key, apiKeyFlag, "a", opts.Auth.Key, "Api key to authorization")

	if smtpProxy := os.Getenv(smtpProxyEnv); smtpProxy != "" {
		opts.Validator.SMTPProxy = smtpProxy
	}

	rootCmd.Flags().StringVar(&opts.Validator.SMTPProxy, smtpProxyFlag, opts.Validator.SMTPProxy, "Proxy for smtp calling")

	if localName := os.Getenv(localNameEnv); localName != "" {
		opts.Validator.LocalName = localName
	}

	rootCmd.Flags().StringVar(&opts.Validator.LocalName, localNameFlag, opts.Validator.LocalName, "LocalName for SMTP HELO command")

	if _, hasFiberStartupMessage := os.LookupEnv(fiberStartupMessageEnv); !hasFiberStartupMessage {
		fiberStartupMessage = hasFiberStartupMessage
	}
	rootCmd.Flags().BoolVar(&fiberStartupMessage, fiberStartupMessageFlag, fiberStartupMessage, "Show or not Fiber startup message")
}
