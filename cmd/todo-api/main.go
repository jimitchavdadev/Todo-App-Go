package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jimitchavdadev/todo-app/internal/config"
	"github.com/jimitchavdadev/todo-app/internal/db"
	"github.com/jimitchavdadev/todo-app/internal/handler"
	"github.com/jimitchavdadev/todo-app/internal/repository"
	"github.com/jimitchavdadev/todo-app/internal/service"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := db.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	apiHandler := handler.NewAPIHandler(taskService)

	router := mux.NewRouter()
	apiHandler.RegisterRoutes(router)

	addr := fmt.Sprintf(":%s", cfg.APIPort)
	log.Printf("Starting API server on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
