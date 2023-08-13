package standard

import (
	"context"
	"log"
	"todo_app_mux/models"
)

func (s *domainService) AddItemToList(ctx context.Context, item string) error {
	err := s.db.AddItemToDb(ctx, item)
	if err != nil {
		log.Println(ctx, err)
		return err
	}

	return nil
}

func (s *domainService) GetTodoItems(ctx context.Context) ([]models.TodoListItem, error) {
	todoItems, err := s.db.GetItemsFromDb(ctx)
	if err != nil {
		log.Println(ctx, err)
		return []models.TodoListItem{}, err
	}

	return todoItems, nil
}

func (s *domainService) UpdateItemFromTodo(ctx context.Context, itemId, item string) error {
	err := s.db.UpdateItemFromDb(ctx, itemId, item)
	if err != nil {
		log.Println(ctx, err)
		return err
	}

	return nil
}

func (s *domainService) DeleteItemFromTodo(ctx context.Context, itemId string) error {
	err := s.db.DeleteItemFromDb(ctx, itemId)
	if err != nil {
		log.Println(ctx, err)
		return err
	}

	return nil
}

func (s *domainService) MarkItemComplete(ctx context.Context, itemId string) error {
	err := s.db.UpdateItemCompletedStatus(ctx, itemId, true)
	if err != nil {
		log.Println(ctx, err)
		return err
	}

	return nil
}

func (s *domainService) MarkItemIncomplete(ctx context.Context, itemId string) error {
	err := s.db.UpdateItemCompletedStatus(ctx, itemId, false)
	if err != nil {
		log.Println(ctx, err)
		return err
	}

	return nil
}
