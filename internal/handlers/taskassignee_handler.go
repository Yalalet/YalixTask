package handlers

import (
	"encoding/json"
	"myapp/internal/models"
	"myapp/internal/service"
	"net/http"
)

type TaskAssigneeHandler struct {
	Service *service.TaskAssigneeService
}

func (h *TaskAssigneeHandler) TaskAssignees(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.createTaskAssignee(w, r)
	case "GET":
		h.listTaskAssignees(w, r)
	}
}

func (h *TaskAssigneeHandler) createTaskAssignee(w http.ResponseWriter, r *http.Request) {
	var assigneeInput struct {
		TaskID     int `json:"task_id"`
		UserID     int `json:"user_id"`
		AssignedBy int `json:"assigned_by"`
	}

	err := json.NewDecoder(r.Body).Decode(&assigneeInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	assignee := &models.TaskAssignee{
		TaskID:     assigneeInput.TaskID,
		UserID:     assigneeInput.UserID,
		AssignedBy: assigneeInput.AssignedBy,
	}

	if err := h.Service.CreateTaskAssignee(assignee); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(assignee)
}

func (h *TaskAssigneeHandler) listTaskAssignees(w http.ResponseWriter, r *http.Request) {
	assignees, err := h.Service.GetAllTaskAssignees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(assignees)
}
