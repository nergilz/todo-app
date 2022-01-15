package main

import (
	todoapp "github.com/nergilz/todo-app"
	"github.com/nergilz/todo-app/pkg/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	handler := new(handler.Handler)
	r := handler.InitHandlers()

	server := new(todoapp.Server)

	if err := server.Run("8080", r); err != nil {
		logrus.Fatal("error running server: %v", err.Error())
	}
}
