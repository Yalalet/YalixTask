package handlers

import (
	"encoding/json"
	"myapp/internal/models"
	"myapp/internal/service"
	"net/http"
	"time"
)

type TaskHandler struct {
	Service *service.TaskService
}

func (h *TaskHandler) Tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.createTask(w, r)
	case "GET":
		h.listTasks(w, r)
	}
}

func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var taskInput struct {
		Name        string    `json:"name"`
		CreatedAt   time.Time `json:"created_at"`
		Deadline    time.Time `json:"deadline"`
		Description string    `json:"description"`
		TeamID      *int      `json:"team_id"`
		StatusID    int       `json:"status_id"`
		PriorityID  int       `json:"priority_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&taskInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := &models.Task{
		Name:        taskInput.Name,
		CreatedAt:   taskInput.CreatedAt,
		Deadline:    taskInput.Deadline,
		Description: taskInput.Description,
		TeamID:      taskInput.TeamID,
		StatusID:    taskInput.StatusID,
		PriorityID:  taskInput.PriorityID,
	}

	if err := h.Service.CreateTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) listTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}
