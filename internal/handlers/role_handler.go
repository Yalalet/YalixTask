package handlers

import (
	"encoding/json"
	"myapp/internal/service"
	"net/http"
)

type RoleHandler struct {
	Service *service.RoleService
}

func (h *RoleHandler) Roles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.listRoles(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *RoleHandler) listRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.Service.GetAllRoles()
	if err != nil {
		http.Error(w, "Failed to fetch roles", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(roles)
}
