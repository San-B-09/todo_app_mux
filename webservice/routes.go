package webservice

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *webService) addBasicRoutes(router *mux.Router) {
	router.HandleFunc("/ping", s.Ping).Methods(http.MethodGet)
}
