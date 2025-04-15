package server

import (
	"api_gateway/internal/category"
	"api_gateway/internal/middleware"
	"api_gateway/internal/routes"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	Category *category.HandlerCategory
}

type Server struct {
	app        *fiber.App
	handlers   Handlers
	cleanupFns []func()
}

func NewServer(handlers Handlers) *Server {
	app := fiber.New(fiber.Config{
		AppName:      "API Gateway",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	})

	middleware.Setup(app)

	setupRoutes(app, handlers)

	return &Server{
		app:      app,
		handlers: handlers,
	}

}

// func (s *Server) AddCleanup(fn func()) {
// 	s.cleanupFns = append(s.cleanupFns, fn)
// }

// func (s *Server) Start() error {
// 	defer func() {
// 		for _, fn := range s.cleanupFns {
// 			fn()
// 		}
// 	}()

// 	return s.app.Listen()
// }

func setupRoutes(app *fiber.App, handlers Handlers) {
	api := app.Group("/api")

	routes.RegisterCategoryRoutes(api, handlers.Category)

	app.Get("/health", healtCheck)

}

func healtCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "OK"})
}

func (s *Server) RegisterCleanup(fn func()) {
	s.cleanupFns = append(s.cleanupFns, fn)
}

func (s *Server) Start(address string) error {
	log.Printf("Starting server on %s", address)
	return s.app.Listen(address)
}

func (s *Server) Shutdown() error {
	defer func() {
		for _, fn := range s.cleanupFns {
			fn()
		}
	}()

	log.Println("Shutting shutdown server gracefully")

	return s.app.Shutdown()
}
