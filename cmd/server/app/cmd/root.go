package cmd

import (
	"fmt"
	v1 "github.com/go-email-validator/go-ev-presenters/pkg/api/v1"
	"github.com/go-email-validator/go-ev-presenters/pkg/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var opts = v1.NewOptions()

var rootCmd = &cobra.Command{
	Use:  "ev",
	Long: "start HTTP server",
	PreRun: func(cmd *cobra.Command, args []string) {
		if isVerbose {
			fiberStartupMessage = true
		}

		opts.Fiber.DisableStartupMessage = !fiberStartupMessage
	},
	Run: func(cmd *cobra.Command, args []string) {
		if isVerbose {
			l, err := zap.NewDevelopment()
			if err != nil {
				panic(err)
			}
			log.SetLogger(l)
		}

		serv := v1.NewServer(opts)

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
