package category

import "github.com/gofiber/fiber/v2"

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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	res, err := h.repo.Create(req.Name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create category",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (h *HandlerCategory) GetCategories(ctx *fiber.Ctx) error {
	res, err := h.repo.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed get categories",
		})
	}

	return ctx.JSON(res)
}

func (h *HandlerCategory) GetCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid Categpry ID",
		})
	}

	res, err := h.repo.GetByID(int32(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "category not found",
		})
	}

	return ctx.JSON(res)
}

func (h *HandlerCategory) UpdateCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Category ID",
		})
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request Body",
		})
	}

	res, err := h.repo.Update(int32(id), req.Name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update category",
		})
	}

	return ctx.JSON(res)
}

func (h *HandlerCategory) DeleteCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"erorr": "invalid category ID",
		})
	}

	res, err := h.repo.Delete(int32(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete category",
		})
	}

	return ctx.JSON(res)
}
