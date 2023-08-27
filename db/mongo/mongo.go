package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"todo_app_mux/db"
	"todo_app_mux/log"
)

type mongoService struct {
	db mongo.Client
}

func New(ctx context.Context, mongoUrl string) db.Idb {
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.GenericError(ctx, errors.New("Error initiating mongo client"))
		return &mongoService{}
	}

	return &mongoService{
		db: *mongoClient,
	}
}
