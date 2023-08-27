package domain

import (
	"context"
	"todo_app_mux/models"
)

type IDomainService interface {
	AddItemToList(ctx context.Context, item string) error
	GetTodoItems(ctx context.Context) ([]models.TodoListItem, error)
	UpdateItemFromTodo(ctx context.Context, itemId, item string) error
	DeleteItemFromTodo(ctx context.Context, itemId string) error
	MarkItemComplete(ctx context.Context, itemId string) error
	MarkItemIncomplete(ctx context.Context, itemId string) error
}
