package cmd

import (
	"fmt"
	"github.com/go-email-validator/go-ev-presenters/pkg/api/v1/server"
	"github.com/go-email-validator/go-ev-presenters/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var opts = server.NewOptions()

var rootCmd = &cobra.Command{
	Use:  "ev",
	Long: `GRPC and HTTP server start`,
	Run: func(cmd *cobra.Command, args []string) {
		if isVerbose {
			logger := logrus.New()
			logger.SetLevel(logrus.DebugLevel)
			log.SetLogger(logger)
		}

		serv := server.NewServer(opts)

		die := make(chan os.Signal, 1)
		signal.Notify(die, os.Interrupt, os.Kill, syscall.SIGTERM)
		go func() {
			<-die
			serv.Shutdown()
		}()

		err := serv.Start()
		if err != nil {
			fmt.Println(err)
		}

		serv.Wait()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
