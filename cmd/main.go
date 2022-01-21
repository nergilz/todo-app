package main

import (
	todoapp "github.com/nergilz/todo-app"
	"github.com/nergilz/todo-app/pkg/handler"
	"github.com/nergilz/todo-app/pkg/repository"
	"github.com/nergilz/todo-app/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)
	router := handler.InitHandlers()

	server := new(todoapp.Server)

	if err := server.Run("8080", router); err != nil {
		logrus.Fatal("error running server: %v", err.Error())
	}
}

// clean architecture (bob mike)
// слои:
// http requests --> handler --> service (бизнес логика) --> repository (работа с DB)
