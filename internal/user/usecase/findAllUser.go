package usecase

import (
	"github.com/andersonsantech/go-learning/internal/user/domain"
)

func (i *UserUsecaseImpl) FindAll() ([]domain.User, error) {
	return i.UserRepository.FindAll()
}