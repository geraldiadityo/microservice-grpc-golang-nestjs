package routes

import (
	"api_gateway/internal/barang"

	"github.com/gofiber/fiber/v2"
)

func RegisterBarangRoutes(router fiber.Router, handler *barang.HandlerBarang) {
	group := router.Group("/barang")
	group.Post("/create-barang", handler.CreateBarang)
	group.Get("/view", handler.GetBarangs)
	group.Get("/view/:id", handler.GetBarang)
	group.Put("/update-barang/:id", handler.UpdateBarang)
	group.Delete("/delete-barang/:id", handler.DeleteBarang)
}
