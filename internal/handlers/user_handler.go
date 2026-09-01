package handlers

import (
	"encoding/json"
	"errors"
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

	user, err := h.Service.Login(loginInput.Login, loginInput.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {

	var userInput struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		//MiddleName   *string    `json:"middle_name,omitempty"`
		Login    string `json:"login"`
		Email    string `json:"email"`
		Password string `json:"password"`
		//Avatar       *string    `json:"avatar,omitempty"`
		//Phone        *string    `json:"phone,omitempty"`
		//GitHub       *string    `json:"github,omitempty"`
		//RoleID    int       `json:"role_id"`
		//IsActive  bool      `json:"is_active"`
		//CreatedAt time.Time `json:"created_at"`
		//UpdatedAt time.Time `json:"updated_at"`
		//DeletedAt    *time.Time `json:"deleted_at,omitempty"`
		//LastLoginAt  *time.Time `json:"last_login_at,omitempty"`
	}

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.Service.Register(userInput.FirstName, userInput.LastName, userInput.Login, userInput.Email, userInput.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
