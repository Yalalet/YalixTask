package handlers

import (
	"encoding/json"
	"log"
	"myapp/internal/service"
	"net/http"
)

type TeamRoleHandler struct {
	Service *service.TeamRoleService
}

func (h *TeamRoleHandler) TeamRoles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.listTeamRoles(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TeamRoleHandler) listTeamRoles(w http.ResponseWriter, r *http.Request) {
	teamroles, err := h.Service.GetAllTeamRoles()
	if err != nil {
		log.Println("GetAllTeamRoles error:", err) // добавь "log" в импорты
		http.Error(w, "Failed to fetch teamroles", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(teamroles)
}
