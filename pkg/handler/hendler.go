package handler

import (
	"github.com/gorilla/mux"
	"github.com/nergilz/todo-app/pkg/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	logrus.Info("Init Handler")

	return &Handler{
		Service: service,
	}
}

func (h *Handler) InitHandlers() *mux.Router {
	muxRouter := mux.NewRouter()

	auth := muxRouter.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/sign_in", h.SignIn)
	auth.HandleFunc("/sign_up", h.SignUp)

	api := muxRouter.PathPrefix("/api").Subrouter()

	lists := api.PathPrefix("/lists").Subrouter()
	lists.HandleFunc("/", h.CreateList).Methods("POST")
	lists.HandleFunc("/", h.GetAllLists).Methods("GET")
	lists.HandleFunc("/:list_id", h.GetListById).Methods("GET")
	lists.HandleFunc("/:list_id", h.UpdateList).Methods("PUT")
	lists.HandleFunc("/:list_id", h.DeleteList).Methods("DELETE")

	items := api.PathPrefix("/items").Subrouter()
	items.HandleFunc("/", h.CreateItem).Methods("POST")
	items.HandleFunc("/", h.GetAllItems).Methods("GET")
	items.HandleFunc("/:item_id", h.GetItemById).Methods("GET")
	items.HandleFunc("/:item_id", h.UpdateItem).Methods("PUT")
	items.HandleFunc("/:item_id", h.DeleteItem).Methods("DELETE")

	logrus.Info("Init Router")

	return muxRouter
}
