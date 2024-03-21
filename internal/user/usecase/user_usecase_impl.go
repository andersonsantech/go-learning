package usecase

import (
	"fmt"

	"github.com/andersonsantech/go-learning/internal/user/domain"
)

type UserInteractor struct {
    UserRepository UserRepository
}

func (i *UserInteractor) RegisterUser(user domain.User) error {
    fmt.Printf("Passando pela minha use case: %v ", user)
    return i.UserRepository.Save(user) 
}

func (i *UserInteractor) FindByID(id int) (domain.User, error) {
    return i.UserRepository.FindByID(id)
}

func (i *UserInteractor) FindAll() ([]domain.User, error) {
    return i.UserRepository.FindAll()
}

func (i *UserInteractor) Delete(id int) error {
    return i.UserRepository.Delete(id)
}