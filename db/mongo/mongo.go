package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"todo_app_mux/db"
)

type mongoService struct {
	db mongo.Client
}

func New(ctx context.Context, mongoUrl string) db.Idb {
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatalf("Error initiating mongo client")
		return &mongoService{}
	}

	return &mongoService{
		db: *mongoClient,
	}
}
