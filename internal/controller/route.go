package controller

import (
	"fartech/wedding-organizer-service/internal/controller/authentication"
	"fartech/wedding-organizer-service/internal/repository/user"
	authUsecase "fartech/wedding-organizer-service/internal/usecase/authentication"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitRoute(app *fiber.App, db *gorm.DB, logger *zap.Logger) {

	userRepository := user.NewUserRepository(db, logger)
	authUsecase := authUsecase.NewAuthenticationUsecase(userRepository, logger)
	authController := authentication.NewAuthenticationController(authUsecase, logger)

	// Auth
	auth := app.Group("/auth")
	auth.Post("/login", authController.Login)
}
