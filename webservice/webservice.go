package webservice

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo_app_mux/domain"
)

type webService struct {
	domain domain.IDomainService
}

func New(domain domain.IDomainService) *webService {
	return &webService{
		domain: domain,
	}
}

func (s *webService) Start(ctx context.Context) {
	router := mux.NewRouter().StrictSlash(true)

	s.addBasicRoutes(router)

	serverHandler := handlers.CORS()(router)
	err := http.ListenAndServe(":8080", serverHandler)
	if err != nil {
		log.Fatalln(ctx, "Error starting the server:", err.Error())
	}

	return
}

func ReturnOKResponse(ctx context.Context, w http.ResponseWriter, data interface{}) {
	response := httpResponse{
		Meta: httpResponseMeta{
			Code:    http.StatusOK,
			Message: "success",
		},
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(response)
	if err != nil {
		log.Fatalln(ctx, err)
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.Fatalln(ctx, err)
	}
}

func ReturnErrorResponse(ctx context.Context, w http.ResponseWriter, errorCode int64, errorMessage string) {
	response := httpResponse{
		Meta: httpResponseMeta{
			Code:    errorCode,
			Message: errorMessage,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(response)
	if err != nil {
		log.Fatalln(ctx, err)
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.Fatalln(ctx, err)
	}
}

func (s *webService) Ping(w http.ResponseWriter, _ *http.Request) {
	ctx := context.TODO()
	log.Println(ctx, "Server Pinged!")
	ReturnOKResponse(ctx, w, "pong")
}
