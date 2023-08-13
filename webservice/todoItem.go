package webservice

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (s *webService) AddItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var item todoItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusBadRequest, "Error decoding request")
		return
	}

	err = s.domain.AddItemToList(ctx, item.Item)
	if err != nil {
		log.Println(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusInternalServerError, "Error adding item to todo list")
		return
	}

	ReturnOKResponse(ctx, w, "success")
	return
}

func (s *webService) GetItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	todoItems, err := s.domain.GetTodoItems(ctx)
	if err != nil {
		log.Println(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusInternalServerError, "Error fetching todo items")
		return
	}

	ReturnOKResponse(ctx, w, todoItems)
	return
}

func (s *webService) UpdateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	itemId := mux.Vars(r)["item-id"]
	if itemId == "" {
		log.Println(ctx, "Empty item id")
		ReturnErrorResponse(ctx, w, http.StatusBadRequest, "Item Id not found")
		return
	}

	var item todoItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusBadRequest, "Error decoding request")
		return
	}

	err = s.domain.UpdateItemFromTodo(ctx, itemId, item.Item)
	if err != nil {
		log.Println(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusInternalServerError, "Error adding item to todo list")
		return
	}

	ReturnOKResponse(ctx, w, "success")
	return
}

func (s *webService) DeleteItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	itemId := mux.Vars(r)["item-id"]
	if itemId == "" {
		log.Println(ctx, "Empty item id")
		ReturnErrorResponse(ctx, w, http.StatusBadRequest, "Item Id not found")
		return
	}

	err := s.domain.DeleteItemFromTodo(ctx, itemId)
	if err != nil {
		log.Println(ctx, err)
		ReturnErrorResponse(ctx, w, http.StatusInternalServerError, "Error adding item to todo list")
		return
	}

	ReturnOKResponse(ctx, w, "success")
	return
}
