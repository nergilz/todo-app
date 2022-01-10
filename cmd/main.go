package main

import (
	todoapp "github.com/nergilz/todo-app"
	"github.com/sirupsen/logrus"
)

func main() {
	server := new(todoapp.Server)
	if err := server.Run("8080"); err != nil {
		logrus.Fatal("error running server: %v", err.Error())
	}
}
