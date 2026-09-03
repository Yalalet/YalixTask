package handlers

import (
	"myapp/internal/service"
	"net/http"
	"strconv"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "неверный формат заголовка Authorization", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := service.ValidateToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		r.Header.Set("user_id", strconv.Itoa(userID)) // сохраняем user_id в заголовках запроса

		next(w, r) // передаём управление дальше, к реальному хендлеру

	}
}

func AdminOnlyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleID := r.Header.Get("role_id") // положили туда в AuthMiddleware
		if roleID != "1" {                // 1 — id роли admin, судя по твоей схеме
			http.Error(w, "доступ запрещён", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
