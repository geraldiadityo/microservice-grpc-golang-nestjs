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
