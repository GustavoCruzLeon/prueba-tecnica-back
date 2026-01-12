package application

import (
	"errors"
	"prueba-tecnica-back/internal/identity/domain"
)

type RegisterUseCase struct {
	userRepo     domain.UserRepository
	passwordHash func(password string) (string, error)
}

func NewRegisterUseCase(userRepo domain.UserRepository, passwordHash func(string) (string, error)) *RegisterUseCase {
	return &RegisterUseCase{
		userRepo:     userRepo,
		passwordHash: passwordHash,
	}
}

func (u *RegisterUseCase) Execute(name, email, password string) (*domain.User, error) {
	if name == "" || email == "" || password == "" {
		return nil, errors.New("name, email and password are required")
	}

	existing, err := u.userRepo.FindByEmail(email)
	if err == nil && existing != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := u.passwordHash(password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	err = u.userRepo.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
