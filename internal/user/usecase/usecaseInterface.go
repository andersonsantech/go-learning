package usecase

import "github.com/andersonsantech/go-learning/internal/user/domain"

type UserUseCase interface {
    RegisterUser(user domain.User) error
    FindByID(id int) (domain.User, error)
    FindAll() ([]domain.User, error)
    Delete(id int) error
}