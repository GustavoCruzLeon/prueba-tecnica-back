package application

import (
	"errors"
	"prueba-tecnica-back/internal/identity/domain"
	"prueba-tecnica-back/internal/identity/infrastructure/jwt"

	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	userRepo   domain.UserRepository
	jwtService *jwt.JWTService
}

func NewLoginUseCase(userRepo domain.UserRepository, jwtService *jwt.JWTService) *LoginUseCase {
	return &LoginUseCase{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (u *LoginUseCase) Execute(email, password string) (*domain.User, string, error) {
	if email == "" || password == "" {
		return nil, "", errors.New("email and password are required")
	}

	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := u.jwtService.GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
