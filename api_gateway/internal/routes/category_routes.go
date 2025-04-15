package routes

import (
	"api_gateway/internal/category"

	"github.com/gofiber/fiber/v2"
)

func RegisterCategoryRoutes(router fiber.Router, handler *category.HandlerCategory) {
	group := router.Group("/category")
	group.Post("/create-category", handler.CreateCategory)
	group.Get("/view", handler.GetCategories)
}
