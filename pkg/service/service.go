package service

import (
	todo "github.com/amenshenin/go1"
	"github.com/amenshenin/go1/pkg/repository"
)

type Autorization interface {
	CreateUser(use todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Autorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
	}
}
