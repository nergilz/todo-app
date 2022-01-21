package repository

import "github.com/sirupsen/logrus"

type Authtorisation interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authtorisation
	TodoList
	TodoItem
}

func NewRepository() *Repository {
	logrus.Info("Init Repository")

	return &Repository{}
}
