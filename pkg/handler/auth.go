package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SignIn"))
	logrus.Info("request sign_in")
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SignUp"))
	logrus.Info("request sign_up")
}
