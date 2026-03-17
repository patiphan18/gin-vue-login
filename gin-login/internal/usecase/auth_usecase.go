package usecase

import (
	"errors"
	"gin-login/internal/delivery/http/dto"
	"gin-login/internal/domain"
	"gin-login/internal/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	userRepo repository.UserRepository
}

func NewAuthUsecase(r repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepo: r}
}

func (u *AuthUsecase) Login(input dto.LoginInput) (*domain.User, error) {
	user, err := u.userRepo.FindByUsername(input.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (u *AuthUsecase) Register(input dto.RegisterInput) error {
	if err := domain.ValidateUsername(input.Username); err != nil {
		return err
	}

	if err := domain.ValidatePassword(input.Password); err != nil {
		return err
	}

	if input.Password != input.ConfirmPassword {
		return errors.New("password and confirm password do not match")
	}

	user, err := u.userRepo.FindByUsername(input.Username)
	if err != nil && err != mongo.ErrNoDocuments {
		return errors.New("failed to find user")
	}

	if user != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return u.userRepo.Create(&domain.User{
		Username: input.Username,
		Password: string(hashedPassword),
	})
}
