package category

import (
	"api_gateway/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type HandlerCategory struct {
	repo RepositoryCategory
}

func NewHandlerCategory(repo RepositoryCategory) *HandlerCategory {
	return &HandlerCategory{
		repo: repo,
	}
}

func (h *HandlerCategory) CreateCategory(ctx *fiber.Ctx) error {
	var req struct {
		Name string `json:"name"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	res, err := h.repo.Create(req.Name)
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusCreated, "Success created category", res)
}

func (h *HandlerCategory) GetCategories(ctx *fiber.Ctx) error {
	res, err := h.repo.GetAll()
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to get categories")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "success", res)
}

func (h *HandlerCategory) GetCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "invalid ID category")
	}

	res, err := h.repo.GetByID(int32(id))
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusNotFound, "category not found!")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "found it", res)
}

func (h *HandlerCategory) UpdateCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid category ID")
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "invalid request body")
	}

	res, err := h.repo.Update(int32(id), req.Name)
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to update category")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "Category success updated", res)
}

func (h *HandlerCategory) DeleteCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "invalid category ID")
	}

	res, err := h.repo.Delete(int32(id))
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to delete category")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "succcess deleted category", res)
}
