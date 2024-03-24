package container

import (
	"github.com/andersonsantech/go-learning/internal/user/delivery/handler"
	"github.com/andersonsantech/go-learning/internal/user/repository"
	"github.com/andersonsantech/go-learning/internal/user/repository/memory"
	"github.com/andersonsantech/go-learning/internal/user/usecase"
)

type Container struct {
	UserRepository repository.UserRepository
	UserUseCase    usecase.UserUseCase
	UserHandler    handler.UserHandler
}

func InjectUserDependencies() *Container {
	userRepositoryImpl := memory.NewUserRepository()
	userUseCaseImpl := &usecase.UserUsecaseImpl{UserRepository: userRepositoryImpl}
	userHandler := &handler.UserHandler{UserUseCase: userUseCaseImpl}

	return &Container{
		UserRepository: userRepositoryImpl,
		UserUseCase:    userUseCaseImpl,
		UserHandler:    *userHandler,
	}
}
