package http

import (
	"errors"
	"fartech/wedding-organizer-service/internal/application_error"

	"github.com/gofiber/fiber/v3"
)

type ApiResponse struct {
	Code    string      `json:"code"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func buildResponse(ctx fiber.Ctx, httpStatus int, responseCode ResponseCode, message ResponseMessage, err ResponseError,
	data interface{},
) error {
	ctx.Status(httpStatus)
	ctx.Response().Header.Set("Content-Type", "application/json")
	return ctx.JSON(ApiResponse{
		Code:    string(responseCode),
		Error:   string(err),
		Message: string(message),
		Data:    data,
	})
}

func BuildSuccessResponse(ctx fiber.Ctx, httpStatus int, responseCode ResponseCode,
	message ResponseMessage, data interface{},
) error {
	return buildResponse(ctx, httpStatus, responseCode, message, "", data)
}

func BuildErrorResponse(ctx fiber.Ctx, err error) error {
	httpStatus := fiber.StatusInternalServerError
	responseCode := InternalServerErrorCode
	responseErr := InternalServerError
	message := InternalServerErrorErrorMessage

	if errors.Is(err, application_error.NotFoundErr) {

		httpStatus = fiber.StatusNotFound
		responseCode = NotFoundErrorCode
		responseErr = ResourceNotFoundError
		message = ResourceNotFoundErrorMessage
	}

	return buildResponse(ctx, httpStatus, responseCode, message, responseErr, nil)
}
