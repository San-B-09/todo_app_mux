package webservice

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"todo_app_mux/log"
	"todo_app_mux/models"
)

func (s *webService) AddItem(w http.ResponseWriter, r *http.Request) {
	ctx := UpgradeContext(r.Context())
	var item models.TodoListItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.GenericError(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusBadRequest, "Error decoding request")
		return
	}

	err = s.domain.AddItemToList(ctx, item.Item)
	if err != nil {
		log.GenericError(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusInternalServerError, "Error adding item to todo list")
		return
	}

	ReturnOKResponse(ctx, w, "success")
	return
}

func (s *webService) GetItems(w http.ResponseWriter, r *http.Request) {
	ctx := UpgradeContext(r.Context())

	todoItems, err := s.domain.GetTodoItems(ctx)
	if err != nil {
		log.GenericError(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusInternalServerError, "Error fetching todo items")
		return
	}

	ReturnOKResponse(ctx, w, todoItems)
	return
}

func (s *webService) UpdateItem(w http.ResponseWriter, r *http.Request) {
	ctx := UpgradeContext(r.Context())
	itemId := mux.Vars(r)["item-id"]
	if itemId == "" {
		log.GenericError(ctx, errors.New("Empty item id"))
		ReturnErrorResponse(ctx, w, http.StatusBadRequest, "Item Id not found")
		return
	}

	var item models.TodoListItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.GenericError(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusBadRequest, "Error decoding request")
		return
	}

	err = s.domain.UpdateItemFromTodo(ctx, itemId, item.Item)
	if err != nil {
		log.GenericError(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusInternalServerError, "Error adding item to todo list")
		return
	}

	ReturnOKResponse(ctx, w, "success")
	return
}

func (s *webService) DeleteItem(w http.ResponseWriter, r *http.Request) {
	ctx := UpgradeContext(r.Context())
	itemId := mux.Vars(r)["item-id"]
	if itemId == "" {
		log.GenericError(ctx, errors.New("Empty item id"))
		ReturnErrorResponse(ctx, w, http.StatusBadRequest, "Item Id not found")
		return
	}

	err := s.domain.DeleteItemFromTodo(ctx, itemId)
	if err != nil {
		log.GenericError(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusInternalServerError, "Error adding item to todo list")
		return
	}

	ReturnOKResponse(ctx, w, "success")
	return
}

func (s *webService) MarkItemComplete(w http.ResponseWriter, r *http.Request) {
	ctx := UpgradeContext(r.Context())
	itemId := mux.Vars(r)["item-id"]
	if itemId == "" {
		log.GenericError(ctx, errors.New("Empty item id"))
		ReturnErrorResponse(ctx, w, http.StatusBadRequest, "Item Id not found")
		return
	}

	err := s.domain.MarkItemComplete(ctx, itemId)
	if err != nil {
		log.GenericError(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusInternalServerError, "Error adding item to todo list")
		return
	}

	ReturnOKResponse(ctx, w, "success")
	return
}

func (s *webService) MarkItemIncomplete(w http.ResponseWriter, r *http.Request) {
	ctx := UpgradeContext(r.Context())
	itemId := mux.Vars(r)["item-id"]
	if itemId == "" {
		log.GenericError(ctx, errors.New("Empty item id"))
		ReturnErrorResponse(ctx, w, http.StatusBadRequest, "Item Id not found")
		return
	}

	err := s.domain.MarkItemIncomplete(ctx, itemId)
	if err != nil {
		log.GenericError(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusInternalServerError, "Error adding item to todo list")
		return
	}

	ReturnOKResponse(ctx, w, "success")
	return
}
