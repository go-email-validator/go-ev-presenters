package v1

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp/smtp_client"
	"github.com/go-email-validator/go-ev-presenters/pkg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
	"h12.io/socks"
	"io/ioutil"
	"net"
	"net/http"
	"net/smtp"
	"strings"
	"sync"
)

func NewServer(opts Options) Server {
	return Server{
		opts: opts,
	}
}

type Server struct {
	httpServer *http.Server
	opts       Options
	waitGroup  sync.WaitGroup
}

func (s *Server) Start() error {
	if err := s.StartHTTP(); err != nil {
		return fmt.Errorf("starting HTTP server failed: %w", err)
	}

	return nil
}

func (s *Server) StartHTTP() error {
	app := fiber.New()
	// Show error openapi format
	app.Use(func(c *fiber.Ctx) error {
		chainErr := c.Next()

		if chainErr != nil {
			return Error(c, chainErr)
		}

		return chainErr
	})
	// logger
	app.Use(func(c *fiber.Ctx) error {
		chainErr := c.Next()

		// TODO Add formatting and fields
		if chainErr != nil {
			log.Logger().Error(chainErr)
		}

		return chainErr
	})
	app.Use(fiberrecover.New())

	var dialFunc evsmtp.DialFunc
	if s.opts.SMTPProxy != "" {
		dialFunc = func(addr string) (smtp_client.SMTPClient, error) {
			conn, err := socks.Dial(s.opts.SMTPProxy)("tcp", addr)
			if err != nil {
				log.Logger().WithFields(logrus.Fields{
					"proxy":   s.opts.SMTPProxy,
					"address": addr,
				}).Errorf("proxy create conn: %s", err)
				return nil, err
			}

			host, _, _ := net.SplitHostPort(addr)
			return smtp.NewClient(conn, host)
		}
	}

	server := defaultInstance(s.opts, dialFunc)
	app.Post("/v1/validation/single", server.EmailValidationSingleValidationPost)
	app.Get("/v1/validation/single/:email", server.EmailValidationSingleValidationGet)

	err := s.addSwagger(app)
	if err != nil {
		return err
	}

	s.waitGroup.Add(1)
	go func() {
		log.Logger().Debugf("HTTP server is starting on %s", s.opts.HTTP.Bind)
		err = app.Listen(s.opts.HTTP.Bind)

		if err != nil && !strings.Contains(err.Error(), "http: Server closed") {
			log.Logger().Errorf("http server: %s", err)
		}
		s.waitGroup.Done()
		log.Logger().Debug("HTTP server stopped")
	}()
	return nil
}

func (s *Server) addSwagger(app *fiber.App) error {
	openapi, err := ioutil.ReadFile(s.opts.HTTP.OpenApiPath)
	if err != nil {
		return err
	}
	app.Get("/swagger.json", func(c *fiber.Ctx) error {
		return c.SendStream(bytes.NewReader(openapi))
	})

	app.Use("/swagger-ui/", filesystem.New(filesystem.Config{
		Root: http.Dir("third_party/swagger-ui"),
	}))

	return nil
}

func (s *Server) Wait() {
	s.waitGroup.Wait()
}

func (s *Server) Shutdown() {
	s.ShutdownHTTP()
}

func (s *Server) ShutdownHTTP() {
	ctx, cancel := context.WithTimeout(context.Background(), s.opts.HTTP.ShutdownTimeout)
	defer cancel()
	err := s.httpServer.Shutdown(ctx)
	log.Logger().Errorf("shutdown http: %v", err)
}
