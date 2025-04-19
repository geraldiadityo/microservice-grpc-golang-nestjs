package routes

import (
	"api_gateway/internal/pengguna"

	"github.com/gofiber/fiber/v2"
)

func RegisterPenggunaRoutes(router fiber.Router, handler *pengguna.HandlerPengguna) {
	group := router.Group("/pengguna")
	group.Post("/create-pengguna", handler.CreatePengguna)
	group.Get("/view", handler.GetPenggunas)
	group.Get("/view/:id", handler.GetPengguna)
	group.Put("/update-pengguna/:id", handler.UpdatePengguna)
	group.Delete("/delete-pengguna/:id", handler.DeletePengguna)
}
