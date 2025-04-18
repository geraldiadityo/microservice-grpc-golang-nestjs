package routes

import (
	"api_gateway/internal/role"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoleRoutes(router fiber.Router, handler *role.HandlerRole) {
	group := router.Group("/role")
	group.Post("/create-role", handler.CreateRole)
	group.Get("/view", handler.GetRoles)
	group.Get("/view/:id", handler.GetRole)
	group.Put("/update-role/:id", handler.UpdateRole)
	group.Delete("/delete-role/:id", handler.DeleteRole)
}
