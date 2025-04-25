package handler

import (
	"fmt"
	"github.com/jimitchavdadev/todo-app/internal/service"
	"strconv"
	"strings"
)

type CLIHandler struct {
	service *service.TaskService
}

func NewCLIHandler(service *service.TaskService) *CLIHandler {
	return &CLIHandler{service: service}
}

func (h *CLIHandler) HandleCommand(input string) {
	args := strings.Fields(input)
	if len(args) == 0 {
		h.printHelp()
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Println("Usage: add <title> [description]")
			return
		}
		desc := ""
		if len(args) > 2 {
			desc = strings.Join(args[2:], " ")
		}
		if err := h.service.CreateTask(args[1], desc); err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println("Task added successfully")
		}

	case "list":
		tasks, err := h.service.ListTasks()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		for _, task := range tasks {
			status := "pending"
			if task.Completed {
				status = "completed"
			}
			fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s, Created: %s\n",
				task.ID, task.Title, task.Description, status, task.CreatedAt)
		}

	case "complete":
		if len(args) < 2 {
			fmt.Println("Usage: complete <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		if err := h.service.CompleteTask(id); err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println("Task marked as complete")
		}

	case "delete":
		if len(args) < 2 {
			fmt.Println("Usage: delete <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		if err := h.service.DeleteTask(id); err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println("Task deleted")
		}

	default:
		h.printHelp()
	}
}

func (h *CLIHandler) printHelp() {
	fmt.Println(`Commands:
  add <title> [description] - Add a new task
  list                     - List all tasks
  complete <id>           - Mark task as complete
  delete <id>             - Delete a task`)
}
