package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jimitchavdadev/todo-app/internal/service"
	"net/http"
	"strconv"
)

type APIHandler struct {
	service *service.TaskService
}

func NewAPIHandler(service *service.TaskService) *APIHandler {
	return &APIHandler{service: service}
}

func (h *APIHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", h.createTask).Methods("POST")
	r.HandleFunc("/tasks", h.listTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}/complete", h.completeTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", h.deleteTask).Methods("DELETE")
}

func (h *APIHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateTask(req.Title, req.Description); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *APIHandler) listTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.ListTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (h *APIHandler) completeTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.service.CompleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *APIHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
