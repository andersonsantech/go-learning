package handler

import "github.com/andersonsantech/go-learning/internal/user/domain"

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
}

type UserResponseReduced struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age: user.Age,
	}
}

func ToUserResponseReduced(user domain.User) UserResponseReduced {
	return UserResponseReduced{
		ID:    user.ID,
		Email: user.Email,
	}
}

func ToUsersResponse(users []domain.User) []UserResponseReduced {
	var result []UserResponseReduced
	for _, user := range users {
		result = append(result, ToUserResponseReduced(user))
	}
	return result
}

func (req *UserRequest) ToUser() domain.User {
	return domain.User{
		Name:  req.Name,
		Email: req.Email,
		Age: req.Age,
	}
}