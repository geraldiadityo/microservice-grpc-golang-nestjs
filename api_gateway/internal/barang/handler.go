package barang

import (
	"api_gateway/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type HandlerBarang struct {
	repo RepositoryBarang
}

func NewHandlerBarang(repo RepositoryBarang) *HandlerBarang {
	return &HandlerBarang{
		repo: repo,
	}
}

func (h *HandlerBarang) CreateBarang(ctx *fiber.Ctx) error {
	var req struct {
		Name       string `json:"name"`
		CategoryId int32  `json:"categoryId"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid body request")
	}

	res, err := h.repo.Create(req.Name, req.CategoryId)
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to create barang")
	}

	return utils.SuccessResponse(ctx, fiber.StatusCreated, "success created barang", res)
}

func (h *HandlerBarang) GetBarangs(ctx *fiber.Ctx) error {
	res, err := h.repo.GetAll()
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to get list barang")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "succecss", res)
}

func (h *HandlerBarang) GetBarang(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid ID barang")
	}

	res, err := h.repo.GetByID(int32(id))
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusNotFound, "barang not found")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "success", res)
}

func (h *HandlerBarang) UpdateBarang(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid barang ID")
	}

	var req struct {
		Name       string `json:"name"`
		CategoryId int32  `json:"categoryId"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	res, err := h.repo.Update(int32(id), req.Name, req.CategoryId)

	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to update barang")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "success updated barang", res)
}

func (h *HandlerBarang) DeleteBarang(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid barang ID")
	}

	res, err := h.repo.Delete(int32(id))
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to deleted barang")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "success deleted barang", res)
}
