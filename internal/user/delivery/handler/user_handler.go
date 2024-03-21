package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/andersonsantech/go-learning/internal/user/usecase"
)

type UserHandler struct {
    UserUseCase usecase.UserUseCase
}

func (h *UserHandler) RegisterUser(response http.ResponseWriter, request *http.Request) {
    var userRequest UserRequest
    json.NewDecoder(request.Body).Decode(&userRequest)

    user := userRequest.ToUser()

    err := h.UserUseCase.RegisterUser(user)
    if err != nil {
        http.Error(response, err.Error(), http.StatusInternalServerError)
        return
    }
    response.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUserByID(response http.ResponseWriter, request *http.Request, id string) {
    num, _ := strconv.Atoi(id)
    user, err := h.UserUseCase.FindByID(num)
    if err != nil {
        http.Error(response, err.Error(), http.StatusNotFound)
        return
    }

    result := ToUserResponse(user)
    json.NewEncoder(response).Encode(result)
}

func (h *UserHandler) FindAllUsers(response http.ResponseWriter, request *http.Request) {
    users, err := h.UserUseCase.FindAll()
    if err != nil {
        http.Error(response, err.Error(), http.StatusInternalServerError)
        return
    }

    response.WriteHeader(http.StatusOK)

    result := ToUsersResponse(users)

    json.NewEncoder(response).Encode(result)
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