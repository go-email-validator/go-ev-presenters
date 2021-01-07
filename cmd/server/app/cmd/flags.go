package cmd

import (
	"github.com/go-email-validator/go-ev-presenters/pkg/api/v1/server"
)

const (
	grpcBindFlag = "grpc-bind"
	httpBindFlag = "http-bind"
	verboseFlag  = "verbose"
)

var isVerbose bool

func init() {
	rootCmd.Flags().StringVar(&opts.GRPC.Bind, grpcBindFlag, server.GRPCDefaultHost, "GRPC bind address")
	rootCmd.Flags().StringVar(&opts.HTTP.Bind, httpBindFlag, server.HTTPDefaultHost, "HTTP bind address")
	rootCmd.Flags().BoolVarP(&isVerbose, verboseFlag, "v", false, "Show DEBUG log information")
}
