package main

import (
	"bufio"
	"fmt"
	"github.com/jimitchavdadev/todo-app/internal/config"
	"github.com/jimitchavdadev/todo-app/internal/db"
	"github.com/jimitchavdadev/todo-app/internal/handler"
	"github.com/jimitchavdadev/todo-app/internal/repository"
	"github.com/jimitchavdadev/todo-app/internal/service"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	db, err := db.NewDB(cfg)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	cliHandler := handler.NewCLIHandler(taskService)

	fmt.Println("Todo List CLI - Type 'help' for commands")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			break
		}
		if input == "help" {
			cliHandler.HandleCommand("help")
			continue
		}
		cliHandler.HandleCommand(input)
	}
}
