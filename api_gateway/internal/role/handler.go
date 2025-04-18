package role

import (
	"api_gateway/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type HandlerRole struct {
	repo RepositoryRole
}

func NewHandlerRole(repo RepositoryRole) *HandlerRole {
	return &HandlerRole{
		repo: repo,
	}
}

func (h *HandlerRole) CreateRole(ctx *fiber.Ctx) error {
	var req struct {
		Name string `json:"name"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	res, err := h.repo.Create(req.Name)
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "failed to create role")
	}

	return utils.SuccessResponse(ctx, fiber.StatusCreated, "Success created data role", res)
}

func (h *HandlerRole) GetRoles(ctx *fiber.Ctx) error {
	res, err := h.repo.GetAll()
	if err != nil {
		utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "failed get data roles")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "success", res)
}

func (h *HandlerRole) GetRole(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid role ID")
	}

	res, err := h.repo.GetByID(int32(id))
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusNotFound, "role ID not found")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "success", res)
}

func (h *HandlerRole) UpdateRole(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid role ID")
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	res, err := h.repo.Update(int32(id), req.Name)
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "failed to update role")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "success updated role data", res)
}

func (h *HandlerRole) DeleteRole(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid role ID")
	}

	res, err := h.repo.Delete(int32(id))
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to delete role data")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "success deleted role data", res)
}
