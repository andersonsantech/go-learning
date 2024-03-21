package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/andersonsantech/go-learning/internal/user/domain"
	"github.com/andersonsantech/go-learning/internal/user/usecase"
)

type UserHandler struct {
    UserUseCase usecase.UserUseCase
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
    var user domain.User
    json.NewDecoder(r.Body).Decode(&user)

    err := h.UserUseCase.RegisterUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request, id string) {
    num, _ := strconv.Atoi(id)
    user, err := h.UserUseCase.FindByID(num)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) FindAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := h.UserUseCase.FindAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, id string) {
    num, _ := strconv.Atoi(id)
    err := h.UserUseCase.Delete(num)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}