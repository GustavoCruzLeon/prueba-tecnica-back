package application

import (
	"errors"
	"prueba-tecnica-back/internal/identity/domain"
)

type LoginUseCase struct {
	userRepo        domain.UserRepository
	passwordCompare func(hashedPassword, password string) bool
	tokenGenerator  func(userID uint) (string, error)
}

func NewLoginUseCase(
	userRepo domain.UserRepository,
	passwordCompare func(string, string) bool,
	tokenGenerator func(uint) (string, error),
) *LoginUseCase {
	return &LoginUseCase{
		userRepo:        userRepo,
		passwordCompare: passwordCompare,
		tokenGenerator:  tokenGenerator,
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

	if !u.passwordCompare(user.Password, password) {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := u.tokenGenerator(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
