package db

import (
	"context"
	"todo_app_mux/models"
)

type ITodoItem interface {
	AddItemToDb(ctx context.Context, item string) error
	GetItemsFromDb(ctx context.Context) ([]models.TodoListItem, error)
	UpdateItemFromDb(ctx context.Context, itemId, item string) error
	DeleteItemFromDb(ctx context.Context, itemId string) error
	UpdateItemCompletedStatus(ctx context.Context, itemId string, itemCompleteStatus bool) error
}
