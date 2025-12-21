package service

import (
	"errors"
	"regexp"

	"prueba-tecnica-back/internal/models"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Register(name, email, password string) (*models.User, error) {
	if err := validatePassword(password); err != nil {
		return nil, err
	}

	var existing models.User
	if err := s.db.Where("email = ?", email).First(&existing).Error; err == nil {
		return nil, errors.New("email ya registrado")
	}

	user := &models.User{
		Name:  name,
		Email: email,
	}

	if err := user.HashPassword(password); err != nil {
		return nil, err
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("credenciales inválidas")
		}
		return nil, err
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("credenciales inválidas")
	}

	return &user, nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("la contraseña debe tener al menos 8 caracteres")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("la contraseña debe contener al menos una mayúscula")
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return errors.New("la contraseña debe contener al menos un número")
	}
	if !regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password) {
		return errors.New("la contraseña debe contener al menos un carácter especial")
	}
	return nil
}
