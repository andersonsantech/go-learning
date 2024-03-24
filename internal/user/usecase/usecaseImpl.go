package usecase

import (
    "github.com/andersonsantech/go-learning/internal/user/repository"
)

type UserUsecaseImpl struct {
	UserRepository repository.UserRepository
}
