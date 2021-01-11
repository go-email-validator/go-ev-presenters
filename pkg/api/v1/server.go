package v1

import (
	"bytes"
	"fmt"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp/smtp_client"
	"github.com/go-email-validator/go-ev-presenters/pkg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
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
	app       *fiber.App
	opts      Options
	waitGroup sync.WaitGroup
}

func (s *Server) Start() error {
	if err := s.StartHTTP(); err != nil {
		return fmt.Errorf("starting HTTP server failed: %w", err)
	}

	return nil
}

func (s *Server) StartHTTP() error {
	s.app = fiber.New(s.opts.Fiber)
	// Show error openapi format
	s.app.Use(func(c *fiber.Ctx) error {
		chainErr := c.Next()

		if chainErr != nil {
			return Error(c, chainErr)
		}

		return chainErr
	})
	// logger
	s.app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if err := log.Logger().Sync(); err != nil {
				fmt.Print(err)
			}
		}()

		chainErr := c.Next()

		// TODO Add formatting and fields
		if chainErr != nil {
			log.Logger().Error(chainErr.Error())
		}

		return chainErr
	})
	s.app.Use(fiberrecover.New())

	var dialFunc evsmtp.DialFunc
	if s.opts.SMTPProxy != "" {
		dialFunc = func(addr string) (smtp_client.SMTPClient, error) {
			conn, err := socks.Dial(s.opts.SMTPProxy)("tcp", addr)
			if err != nil {
				log.Logger().Error(fmt.Sprintf("proxy create conn: %s", err),
					zap.String("proxy", s.opts.SMTPProxy),
					zap.String("address", addr),
				)
				return nil, err
			}

			host, _, _ := net.SplitHostPort(addr)
			return smtp.NewClient(conn, host)
		}
	}

	server := defaultInstance(s.opts, dialFunc)
	s.app.Post("/v1/validation/single", server.EmailValidationSingleValidationPost)
	s.app.Get("/v1/validation/single/:email", server.EmailValidationSingleValidationGet)

	err := s.addSwagger()
	if err != nil {
		return err
	}

	s.waitGroup.Add(1)
	go func() {
		log.Logger().Debug(fmt.Sprintf("HTTP server is starting on %s", s.opts.HTTP.Bind))
		err = s.app.Listen(s.opts.HTTP.Bind)

		if err != nil && !strings.Contains(err.Error(), "http: Server closed") {
			log.Logger().Error(fmt.Sprintf("http server: %s", err))
		}
		s.waitGroup.Done()
		log.Logger().Debug("HTTP server stopped")
	}()
	return nil
}

func (s *Server) addSwagger() error {
	openapi, err := ioutil.ReadFile(s.opts.HTTP.OpenApiPath)
	if err != nil {
		return err
	}
	s.app.Get("/swagger.json", func(c *fiber.Ctx) error {
		return c.SendStream(bytes.NewReader(openapi))
	})

	s.app.Use("/swagger-ui/", filesystem.New(filesystem.Config{
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
	err := s.app.Shutdown()
	log.Logger().Error(fmt.Sprintf("shutdown http: %v", err))
}
