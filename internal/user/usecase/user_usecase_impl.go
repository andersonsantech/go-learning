package usecase

import (
	"fmt"

	"github.com/andersonsantech/go-learning/internal/user/domain"
)

type UserUsecaseImpl struct {
	UserRepository UserRepository
}

func (i *UserUsecaseImpl) RegisterUser(user domain.User) error {
	fmt.Printf("Passando pela minha use case: %v ", user)
	return i.UserRepository.Save(user)
}

func (i *UserUsecaseImpl) FindByID(id int) (domain.User, error) {
	return i.UserRepository.FindByID(id)
}

func (i *UserUsecaseImpl) FindAll() ([]domain.User, error) {
	return i.UserRepository.FindAll()
}

func (i *UserUsecaseImpl) Delete(id int) error {
	return i.UserRepository.Delete(id)
}
