package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"todo_app_mux/db/mongo"
	"todo_app_mux/domain/standard"
	"todo_app_mux/webservice"
)

func main() {
	ctx := context.Background()

	mongoService := mongo.New(ctx, "mongodb://localhost:27017/")
	domainService := standard.New(mongoService)
	webService := webservice.New(domainService)
	webService.Start(ctx)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
	// Stop and close the connections
	log.Println(ctx, "Stopping Server ")
}
