package usecase

import (
	"fmt"
	"github.com/andersonsantech/go-learning/internal/user/domain"
)

func (i *UserUsecaseImpl) RegisterUser(user domain.User) error {
	fmt.Printf("Passando pela minha use case: %v ", user)
	return i.UserRepository.Save(user)
}