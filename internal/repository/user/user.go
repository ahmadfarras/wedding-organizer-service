package user

import (
	"fartech/wedding-organizer-service/internal/model/domain"
	"fartech/wedding-organizer-service/internal/repository"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

// GetUserByEmail implements repository.UserRepository.
func (u *UserRepository) GetUserByEmail(email string) (user *domain.User, err error) {
	result := u.db.Take(&user, "email = ?", email)
	if result.Error != nil {
		u.logger.Error("failed to get user by email", zap.String("email", email), zap.Error(result.Error))
		return nil, result.Error
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB, logger *zap.Logger) repository.UserRepository {
	return &UserRepository{
		db,
		logger,
	}
}
