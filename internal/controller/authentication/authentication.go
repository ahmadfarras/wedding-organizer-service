package authentication

import (
	"fartech/wedding-organizer-service/internal/model/request"
	"fartech/wedding-organizer-service/internal/usecase"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type AuthenticationController struct {
	authenticationUsecase usecase.AuthenticationUsecase
	logger                *zap.Logger
}

func NewAuthenticationController(
	authenticationUsecase usecase.AuthenticationUsecase,
	logger *zap.Logger,
) *AuthenticationController {
	return &AuthenticationController{
		authenticationUsecase,
		logger,
	}
}

func (a *AuthenticationController) Login(ctx fiber.Ctx) error {
	var req request.LoginRequest
	if err := ctx.Bind().JSON(&req); err != nil {
		a.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resp, err := a.authenticationUsecase.Login(ctx.Context(), req)
	if err != nil {
		a.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
