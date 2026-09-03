package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"myapp/internal/models"
	"myapp/internal/service"
	"net/http"
)

type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.createUser(w, r)
	case "GET":
		h.listUsers(w, r)
	default:
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.loginUser(w, r)
	default:
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) loginUser(w http.ResponseWriter, r *http.Request) {
	var loginInput struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginInput)
	if err != nil {
		http.Error(w, "не удалось прочитать данные", http.StatusBadRequest)
		return
	}

	user, token, err := h.Service.Login(loginInput.Login, loginInput.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	loginResponse := struct {
		User  *models.User `json:"user"`
		Token string       `json:"token"`
	}{
		User:  user,
		Token: token,
	}

	json.NewEncoder(w).Encode(loginResponse)
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {

	var userInput struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Login     string `json:"login"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.Service.Register(userInput.FirstName, userInput.LastName, userInput.Login, userInput.Email, userInput.Password)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrMissingFields):
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case errors.Is(err, service.ErrInvalidEmail):
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case errors.Is(err, service.ErrInvalidPassword):
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case errors.Is(err, service.ErrXSS):
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			log.Printf("Ошибка при создании пользователя: %v", err)
			http.Error(w, service.ErrServerError.Error(), http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAll()

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := range users {
		users[i].FirstName = inputSanitization(users[i].FirstName)
		users[i].LastName = inputSanitization(users[i].LastName)
		users[i].Login = inputSanitization(users[i].Login)
	}

	json.NewEncoder(w).Encode(users)
}
