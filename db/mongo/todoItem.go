package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todo_app_mux/log"
	"todo_app_mux/models"
)

func (m *mongoService) AddItemToDb(ctx context.Context, item string) error {
	insertObject := models.TodoListItem{
		Item: item,
	}

	collection := m.db.Database(defaultDb).Collection(todoListCollection)
	_, err := collection.InsertOne(ctx, insertObject)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoService) GetItemsFromDb(ctx context.Context) ([]models.TodoListItem, error) {
	collection := m.db.Database(defaultDb).Collection(todoListCollection)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []models.TodoListItem{}, err
	}

	defer cursor.Close(ctx)
	var todoItems []models.TodoListItem

	err = cursor.All(ctx, &todoItems)
	if err != nil {
		return []models.TodoListItem{}, err
	}

	return todoItems, nil
}

func (m *mongoService) UpdateItemFromDb(ctx context.Context, itemId, item string) error {
	objId, objectIdErr := primitive.ObjectIDFromHex(itemId)
	if objectIdErr != nil {
		log.GenericError(ctx, objectIdErr)
		return objectIdErr
	}
	updateObject := models.TodoListItem{
		Item: item,
	}

	collection := m.db.Database(defaultDb).Collection(todoListCollection)
	_, err := collection.UpdateByID(ctx, objId, bson.M{"$set": updateObject})
	if err != nil {
		log.GenericError(ctx, err)
		return err
	}

	return nil
}

func (m *mongoService) DeleteItemFromDb(ctx context.Context, itemId string) error {
	objId, objectIdErr := primitive.ObjectIDFromHex(itemId)
	if objectIdErr != nil {
		log.GenericError(ctx, objectIdErr)
		return objectIdErr
	}
	deleteFilter := bson.M{
		"_id": objId,
	}

	collection := m.db.Database(defaultDb).Collection(todoListCollection)
	_, err := collection.DeleteOne(ctx, deleteFilter)
	if err != nil {
		log.GenericError(ctx, err)
		return err
	}

	return nil
}

func (m *mongoService) UpdateItemCompletedStatus(ctx context.Context, itemId string, itemCompleteStatus bool) error {
	objId, objectIdErr := primitive.ObjectIDFromHex(itemId)
	if objectIdErr != nil {
		log.GenericError(ctx, objectIdErr)
		return objectIdErr
	}
	updateObject := bson.M{
		"isCompleted": itemCompleteStatus,
	}

	collection := m.db.Database(defaultDb).Collection(todoListCollection)
	_, err := collection.UpdateByID(ctx, objId, bson.M{"$set": updateObject})
	if err != nil {
		log.GenericError(ctx, err)
		return err
	}

	return nil
}
