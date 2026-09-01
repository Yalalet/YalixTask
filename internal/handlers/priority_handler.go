package handlers

import (
	"encoding/json"
	"myapp/internal/service"
	"net/http"
)

type PriorityHandler struct {
	Service *service.PriorityService
}

func (h *PriorityHandler) Prioritys(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.listPriority(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *PriorityHandler) listPriority(w http.ResponseWriter, r *http.Request) {
	roles, err := h.Service.GetAllPriority()
	if err != nil {
		http.Error(w, "Failed to fetch roles", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(roles)
}
