package application

import (
	"errors"
	"prueba-tecnica-back/internal/identity/domain"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUseCase struct {
	userRepo domain.UserRepository
}

func NewRegisterUseCase(userRepo domain.UserRepository) *RegisterUseCase {
	return &RegisterUseCase{
		userRepo: userRepo,
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = u.userRepo.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
