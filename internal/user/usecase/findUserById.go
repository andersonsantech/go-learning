package usecase

import (
	"github.com/andersonsantech/go-learning/internal/user/domain"
)

func (i *UserUsecaseImpl) FindByID(id int) (domain.User, error) {
	return i.UserRepository.FindByID(id)
}
