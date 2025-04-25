package service

import (
	"github.com/jimitchavdadev/todo-app/internal/models"
	"github.com/jimitchavdadev/todo-app/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title, description string) error {
	task := &models.Task{
		Title:       title,
		Description: description,
		Completed:   false,
	}
	return s.repo.Create(task)
}

func (s *TaskService) ListTasks() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) CompleteTask(id int) error {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	task.Completed = true
	return s.repo.Update(task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
