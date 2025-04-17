package utils

import "github.com/gofiber/fiber/v2"

type SuccessedResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorsResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SuccessResponse(ctx *fiber.Ctx, status int, message string, data interface{}) error {
	return ctx.Status(status).JSON(SuccessedResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(ErrorsResponse{
		Status:  "error",
		Message: message,
	})
}
