package repository

import "fartech/wedding-organizer-service/internal/model/domain"

type UserRepository interface {
	GetUserByEmail(email string) (*domain.User, error)
}
