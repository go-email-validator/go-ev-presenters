package cmd

import (
	"os"
)

const (
	envPrefix     = "EV_"
	httpBindFlag  = "http-bind"
	httpBindEnv   = envPrefix + "HTTP_BIND"
	verboseFlag   = "verbose"
	verboseEnv    = envPrefix + "VERBOSE"
	apiKeyFlag    = "api-key"
	apiKeyEnv     = envPrefix + "API_KEY"
	smtpProxyFlag = "smtp-proxy"
	smtpProxyEnv  = envPrefix + "SMTP_PROXY"
)

var isVerbose bool

func init() {
	// For Heroku
	if gatewayPort := os.Getenv("PORT"); gatewayPort != "" {
		opts.HTTP.Bind = "0.0.0.0:" + gatewayPort
	}

	if bind := os.Getenv(httpBindEnv); bind != "" {
		opts.HTTP.Bind = bind
	}

	_, isVerbose = os.LookupEnv(verboseEnv)

	if apiKey := os.Getenv(apiKeyEnv); apiKey != "" {
		opts.Auth.Key = apiKey
	}

	if smtpProxy := os.Getenv(smtpProxyEnv); smtpProxy != "" {
		opts.SMTPProxy = smtpProxy
	}

	rootCmd.Flags().StringVar(&opts.HTTP.Bind, httpBindFlag, opts.HTTP.Bind, "HTTP bind address")
	rootCmd.Flags().BoolVarP(&isVerbose, verboseFlag, "v", isVerbose, "Show DEBUG log information")

	rootCmd.Flags().StringVarP(&opts.Auth.Key, apiKeyFlag, "a", opts.Auth.Key, "Api key to authorization")

	rootCmd.Flags().StringVar(&opts.SMTPProxy, smtpProxyFlag, opts.SMTPProxy, "Proxy for smtp calling")
}
