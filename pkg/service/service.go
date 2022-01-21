package service

import (
	"github.com/nergilz/todo-app/pkg/repository"
	"github.com/sirupsen/logrus"
)

type Authtorisation interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authtorisation
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	logrus.Info("Init Service")

	return &Service{}
}
