package server

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	fiber *fiber.App
}

func NewServer() *Server {
	app := fiber.New()

	return &Server{
		fiber: app,
	}
}

func (s *Server) Init(address string) {
	s.routeConfig()

	err := s.fiber.Listen(address)
	if err != nil {
		log.Fatalf("Failed To Start The Server: %v", err)
	}
}

func (s *Server) routeConfig() {
	apis := s.fiber.Group("/apis")

	apis.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World")
	})
}
