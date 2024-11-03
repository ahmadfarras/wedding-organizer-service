package usecase

import (
	"context"
	"fartech/wedding-organizer-service/internal/model/request"
	"fartech/wedding-organizer-service/internal/model/response"
)

type AuthenticationUsecase interface {
	Login(ctx context.Context, req request.LoginRequest) (*response.LoginResponse, error)
}
