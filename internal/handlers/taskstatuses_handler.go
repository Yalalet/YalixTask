package handlers

import (
	"encoding/json"
	"myapp/internal/service"
	"net/http"
)

type TaskStatusesHandler struct {
	Service *service.TaskStatusesService
}

func (h *TaskStatusesHandler) TaskStatuses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.listTaskStatuses(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (h *TaskStatusesHandler) listTaskStatuses(w http.ResponseWriter, r *http.Request) {
	taskStatuses, err := h.Service.GetAllTaskStatuses()
	if err != nil {
		http.Error(w, "Failed to fetch task statuses", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(taskStatuses)
}
