package handlers

import (
	"encoding/json"
	"myapp/internal/models"
	"myapp/internal/service"
	"net/http"
	"strconv"
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
	case "PUT":
		h.updateTask(w, r)
	case "DELETE":
		h.deleteTask(w, r)
	default:
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

// =====================================POST==================================
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

// =============================================================================

// =====================================DELETE==================================
func (h *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Query().Get("id")
	if taskID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	taskIDInt, err := strconv.Atoi(taskID)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	err = h.Service.DeleteTask(taskIDInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//==============================================================================

// =====================================UPDATE==================================
func (h *TaskHandler) updateTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Query().Get("id")
	if taskID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

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

	err = h.Service.UpdateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

//==============================================================================

// =====================================GET==================================

func (h *TaskHandler) listTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

// =============================================================================
