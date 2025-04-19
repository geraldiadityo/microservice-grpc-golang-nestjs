package pengguna

import (
	"api_gateway/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type HandlerPengguna struct {
	repo PenggunaRepository
}

func NewHandlerPengguna(repo PenggunaRepository) *HandlerPengguna {
	return &HandlerPengguna{
		repo: repo,
	}
}

func (h *HandlerPengguna) CreatePengguna(ctx *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Password string `json:"password"`
		RoleId   int32  `json:"roleId"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	res, err := h.repo.Create(req.Username, req.Name, req.Password, req.RoleId)
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to created pengguna")
	}

	return utils.SuccessResponse(ctx, fiber.StatusCreated, "Success created data pengguna", res)
}

func (h *HandlerPengguna) GetPenggunas(ctx *fiber.Ctx) error {
	res, err := h.repo.GetAll()
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed get data list pengguna")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "success", res)
}

func (h *HandlerPengguna) GetPengguna(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid pengguna id")
	}

	res, err := h.repo.GetByID(int32(id))
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusNotFound, "pengguna not found!")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "success", res)
}

func (h *HandlerPengguna) UpdatePengguna(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid pengguna ID")
	}

	var req struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Password string `json:"password"`
		RoleId   int32  `json:"roleId"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	res, err := h.repo.Update(int32(id), req.Username, req.Name, req.Password, req.RoleId)
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "failed to update pengguna")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "success updated data pengguna", res)
}

func (h *HandlerPengguna) DeletePengguna(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid pengguna ID")
	}

	res, err := h.repo.Delete(int32(id))
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to delete pengguna")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "success deleted data pengguna", res)
}
