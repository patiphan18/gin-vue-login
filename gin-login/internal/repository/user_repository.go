package repository

import "gin-login/internal/domain"

type UserRepository interface {
	FindByUsername(username string) (*domain.User, error)
	Create(user *domain.User) error
}
