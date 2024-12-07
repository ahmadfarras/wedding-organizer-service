package authentication

import (
	"context"
	"fartech/wedding-organizer-service/internal/application_error"
	"fartech/wedding-organizer-service/internal/model/request"
	"fartech/wedding-organizer-service/internal/model/response"
	"fartech/wedding-organizer-service/internal/repository"
	"fartech/wedding-organizer-service/internal/usecase"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type authenticationUsecase struct {
	userRepository repository.UserRepository
	logger         *zap.Logger
}

// Login implements usecase.AuthenticationUsecase.
func (a *authenticationUsecase) Login(ctx context.Context, req request.LoginRequest,
) (resp *response.LoginResponse, err error) {
	user, err := a.userRepository.GetUserByEmail(req.Email)
	if err != nil {
		a.logger.Error(err.Error())
		return nil, err
	}

	if user == nil {
		return nil, application_error.NotFoundErr
	}

	token, err := uuid.NewV7()
	if err != nil {
		a.logger.Error(err.Error())
		return nil, err
	}

	return response.BuildLoginResponse(token.String()), nil
}

func NewAuthenticationUsecase(userRepository repository.UserRepository, logger *zap.Logger,
) usecase.AuthenticationUsecase {
	return &authenticationUsecase{
		userRepository,
		logger,
	}
}
