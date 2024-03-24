package memory

import (
	"errors"
	"fmt"
	"strconv"
	"github.com/andersonsantech/go-learning/internal/user/domain"
)

type UserRepository struct {
    users map[int]domain.User
}

func NewUserRepository() *UserRepository {
    return &UserRepository{
        users: make(map[int]domain.User),
    }
}

func (r *UserRepository) Save(user domain.User) error {
   fmt.Printf("Creating new UserRepository\n")
    user.ID = len(r.users) + 1
    r.users[user.ID] = user

    fmt.Printf("USers: %v\n", r.users)
    return nil
}

func (r *UserRepository) FindByID(id int) (domain.User, error) {
    user, ok := r.users[id]
    fmt.Printf("User: %v\n", user)
    if !ok {
        return domain.User{}, errors.New("User with ID not found " + strconv.Itoa(id))
    }
    return user, nil
}

func (r *UserRepository) FindAll() ([]domain.User, error) {
    var users []domain.User
    for _, user := range r.users {
        users = append(users, user)
    }
    return users, nil
}

func (r *UserRepository) Delete(id int) error {
    delete(r.users, id)
    return nil
}