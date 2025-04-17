package routes

import (
	"api_gateway/internal/barang"

	"github.com/gofiber/fiber/v2"
)

func RegisterBarangRoutes(router fiber.Router, handler *barang.HandlerBarang) {
	group := router.Group("/barang")
	group.Post("/create-barang", handler.CreateBarang)
}
