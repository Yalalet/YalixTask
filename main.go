package main

import (
	"fmt"
	"myapp/internal/database"
	"myapp/internal/handlers"
	"myapp/internal/repository"
	"myapp/internal/service"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var allowedOrigins = map[string]bool{
	"http://127.0.0.1:5500":     true,
	"http://192.168.0.101:5500": true,
	"http://localhost:5500":     true,
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	dsn := os.Getenv("DATABASE_URL")
	db, err := database.Connect(dsn)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	taskStatusRepo := &repository.TaskStatusesRepository{DB: db}
	taskStatusService := &service.TaskStatusesService{Repo: taskStatusRepo}
	taskStatusHandler := &handlers.TaskStatusesHandler{Service: taskStatusService}

	priorityRepo := &repository.PriorityRepository{DB: db}
	priorityService := &service.PriorityService{Repo: priorityRepo}
	priorityHandler := &handlers.PriorityHandler{Service: priorityService}

	userRepo := &repository.UserRepository{DB: db}
	userService := &service.UserService{Repo: userRepo}
	userHandler := &handlers.UserHandler{Service: userService}

	roleRepo := &repository.RoleRepository{DB: db}
	roleService := &service.RoleService{Repo: roleRepo}
	roleHandler := &handlers.RoleHandler{Service: roleService}

	teamRepo := &repository.TeamRepository{DB: db}
	teamService := &service.TeamService{Repo: teamRepo}
	teamHandler := &handlers.TeamHandler{Service: teamService}

	teamRoleRepo := &repository.TeamRoleRepository{DB: db}
	teamRoleService := &service.TeamRoleService{Repo: teamRoleRepo}
	teamRoleHandler := &handlers.TeamRoleHandler{Service: teamRoleService}

	taskRepo := &repository.TaskRepository{DB: db}
	taskService := &service.TaskService{Repo: taskRepo}
	taskHandler := &handlers.TaskHandler{Service: taskService}

	taskAssigneeRepo := &repository.TaskAssigneeRepository{DB: db}
	taskAssigneeService := &service.TaskAssigneeService{Repo: taskAssigneeRepo}
	taskAssigneeHandler := &handlers.TaskAssigneeHandler{Service: taskAssigneeService}

	teamUserRepo := &repository.TeamUserRepository{DB: db}
	teamUserService := &service.TeamUserService{Repo: teamUserRepo}
	teamUserHandler := &handlers.TeamUserHandler{Service: teamUserService}

	mux := http.NewServeMux()
	mux.HandleFunc("/users", handlers.AuthMiddleware(userHandler.Users))
	mux.HandleFunc("/login", userHandler.Login)
	mux.HandleFunc("/roles", handlers.AuthMiddleware(roleHandler.Roles))
	mux.HandleFunc("/prioritys", handlers.AdminOnlyMiddleware(handlers.AuthMiddleware(priorityHandler.Prioritys)))
	mux.HandleFunc("/teamroles", handlers.AdminOnlyMiddleware(handlers.AuthMiddleware(teamRoleHandler.TeamRoles)))
	mux.HandleFunc("/taskstatuses", taskStatusHandler.TaskStatuses)
	mux.HandleFunc("/teams", handlers.AuthMiddleware(teamHandler.Teams))
	mux.HandleFunc("/tasks", handlers.AuthMiddleware(taskHandler.Tasks))
	mux.HandleFunc("/taskassignees", taskAssigneeHandler.TaskAssignees)
	mux.HandleFunc("/teamuser", teamUserHandler.TeamUsers)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", enableCORS(mux))
}
