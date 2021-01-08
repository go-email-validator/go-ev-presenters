package cmd

import (
	"github.com/go-email-validator/go-ev-presenters/pkg/api/v1/server"
	"os"
)

const (
	envPrefix     = "EV_"
	grpcBindFlag  = "grpc-bind"
	grpcBindEnv   = envPrefix + "GRPC_BIND"
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

	if bind := os.Getenv(grpcBindEnv); bind != "" {
		opts.GRPC.Bind = bind
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

	rootCmd.Flags().StringVar(&opts.GRPC.Bind, grpcBindFlag, server.GRPCDefaultHost, "GRPC bind address")
	rootCmd.Flags().StringVar(&opts.HTTP.Bind, httpBindFlag, opts.HTTP.Bind, "HTTP bind address")
	rootCmd.Flags().BoolVarP(&isVerbose, verboseFlag, "v", isVerbose, "Show DEBUG log information")

	rootCmd.Flags().StringVarP(&opts.Auth.Key, apiKeyFlag, "a", opts.Auth.Key, "Api key to authorization")

	rootCmd.Flags().StringVar(&opts.SMTPProxy, smtpProxyFlag, opts.SMTPProxy, "Proxy for smtp calling")
}
