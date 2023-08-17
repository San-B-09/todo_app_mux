package webservice

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"todo_app_mux/domain"
	"todo_app_mux/log"
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
		log.GenericError(ctx, err, map[string]interface{}{"msg": "Error starting the server:"})
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
	ctxRequestId, _ := ctx.Value(requestId).(string)
	w.Header().Set("requestId", ctxRequestId)
	w.WriteHeader(http.StatusOK)
	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(response)
	if err != nil {
		log.GenericError(ctx, err)
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.GenericError(ctx, err)
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
	ctxRequestId, _ := ctx.Value(requestId).(string)
	w.Header().Set("requestId", ctxRequestId)
	w.WriteHeader(http.StatusOK)
	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(response)
	if err != nil {
		log.GenericError(ctx, err)
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.GenericError(ctx, err)
	}
}

func UpgradeContext(ctx context.Context) context.Context {
	newRequestID := strings.Replace(uuid.New().String(), "-", "", -1)
	return context.WithValue(ctx, requestId, newRequestID)
}

func (s *webService) Ping(w http.ResponseWriter, r *http.Request) {
	ctx := UpgradeContext(r.Context())
	log.GenericInfo(ctx, "Server Pinged")
	ReturnOKResponse(ctx, w, "pong")
}
