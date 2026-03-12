package grades

import (
	"distributed/pkg/password"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Login(username, rawPassword string) (*UserModel, error) {
	var user UserModel
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	if !password.CheckPassword(user.Password, rawPassword) {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}
