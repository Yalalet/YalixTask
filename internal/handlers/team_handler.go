package handlers

import (
	"encoding/json"
	"myapp/internal/service"
	"net/http"
)

type TeamHandler struct {
	Service *service.TeamService
}

func (h *TeamHandler) Teams(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.listTeams(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TeamHandler) listTeams(w http.ResponseWriter, r *http.Request) {
	roles, err := h.Service.GetAllTeams()
	if err != nil {
		http.Error(w, "Failed to fetch roles", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(roles)
}
