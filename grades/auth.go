package grades

import "gorm.io/gorm"

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Login(username, password string) (*UserModel, error) {
	var user UserModel
	if err := s.db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
