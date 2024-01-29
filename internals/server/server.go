package server

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

type Server struct {
	fib *fiber.App
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "App v1.0",
	})

	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Bangkok",
	}))

	return &Server{
		fib: app,
	}
}

func (s *Server) Init(address string) {
	s.routeConfig()

	err := s.fib.Listen(address)
	if err != nil {
		log.Fatalf("Failed To Start The Server: %v", err)
	}
}

func (s *Server) routeConfig() {
	apis := s.fib.Group("/apis")

	// apis.Get("/monitor", monitor.New(monitor.Config{Title: "Monitor Page"}))
	apis.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World")
	})
}
