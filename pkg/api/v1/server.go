package v1

import (
	"bytes"
	"fmt"
	"github.com/go-email-validator/go-ev-presenters/pkg/log"
	"github.com/go-email-validator/go-ev-presenters/statik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"sync"
)

func NewServer(fiberFactory FiberFactory, opts Options) Server {
	if fiberFactory == nil {
		fiberFactory = DefaultFiberFactory
	}

	if opts.IsVerbose {
		opts.Fiber.DisableStartupMessage = false
	}

	if opts.IsVerbose {
		l, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		log.SetLogger(l)
	}

	return Server{
		fiberFactory: fiberFactory,
		opts:         opts,
	}
}

type Server struct {
	fiberFactory FiberFactory
	app          *fiber.App
	opts         Options
	waitGroup    sync.WaitGroup
}

func (s *Server) Start() error {
	if err := s.StartHTTP(); err != nil {
		return fmt.Errorf("starting HTTP server failed: %w", err)
	}

	return nil
}

func (s *Server) StartHTTP() error {
	s.app = s.fiberFactory(DefaultInstance(s.opts), s.opts)

	err := s.addSwagger()
	if err != nil {
		return err
	}

	s.waitGroup.Add(1)
	go func() {
		defer func() {
			if err := log.Logger().Sync(); err != nil {
				fmt.Print(err)
			}
		}()
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
	if !s.opts.HTTP.ShowOpenApi {
		return nil
	}

	openapi, err := statik.ReadFile(s.opts.HTTP.OpenApiPath)
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
