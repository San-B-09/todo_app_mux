package webservice

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *webService) addBasicRoutes(router *mux.Router) {
	router.HandleFunc("/ping", s.Ping).Methods(http.MethodGet)

	router.HandleFunc("/item", s.AddItem).Methods(http.MethodPost)
	router.HandleFunc("/items", s.GetItems).Methods(http.MethodGet)
	router.HandleFunc("/item/{item-id}", s.UpdateItem).Methods(http.MethodPut)
	router.HandleFunc("/item/{item-id}", s.DeleteItem).Methods(http.MethodDelete)

	router.HandleFunc("/item/mark-complete/{item-id}", s.MarkItemComplete).Methods(http.MethodPut)
	router.HandleFunc("/item/mark-incomplete/{item-id}", s.MarkItemIncomplete).Methods(http.MethodPut)
}
