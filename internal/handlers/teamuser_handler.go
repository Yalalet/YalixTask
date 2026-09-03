package handlers

import (
	"encoding/json"
	"myapp/internal/models"
	"myapp/internal/service"
	"net/http"
)

type TeamUserHandler struct {
	Service *service.TeamUserService
}

func (h *TeamUserHandler) TeamUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.createTeamUser(w, r)
	case "GET":
		h.listTeamUsers(w, r)
	default:
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func (h *TeamUserHandler) createTeamUser(w http.ResponseWriter, r *http.Request) {
	var teamUserInput struct {
		TeamID     int `json:"team_id"`
		UserID     int `json:"user_id"`
		TeamRoleID int `json:"team_role_id"`
		InvitedBy  int `json:"invited_by"`
	}

	err := json.NewDecoder(r.Body).Decode(&teamUserInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	teamUser := &models.TeamUser{
		TeamID:     teamUserInput.TeamID,
		UserID:     teamUserInput.UserID,
		TeamRoleID: teamUserInput.TeamRoleID,
		InvitedBy:  teamUserInput.InvitedBy,
	}

	if err := h.Service.CreateTeamUser(teamUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(teamUser)
}

func (h *TeamUserHandler) listTeamUsers(w http.ResponseWriter, r *http.Request) {
	teamUsers, err := h.Service.GetAllTeamUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(teamUsers)
}
