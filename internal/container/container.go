package container

import (
	"github.com/andersonsantech/go-learning/internal/user/delivery/handler"
	"github.com/andersonsantech/go-learning/internal/user/repository/memory"
	"github.com/andersonsantech/go-learning/internal/user/usecase"
)

type Container struct {
	UserRepository usecase.UserRepository
	UserUseCase    usecase.UserUseCase
	UserHandler    handler.UserHandler
}

func InjectUserDependencies() *Container {
	userRepository := memory.NewUserRepository()
	userUseCase := &usecase.UserInteractor{UserRepository: userRepository}
	userHandler := &handler.UserHandler{UserUseCase: userUseCase}

	return &Container{
		UserRepository: userRepository,
		UserUseCase:    userUseCase,
		UserHandler:    *userHandler,
	}
}
