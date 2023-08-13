package models

type TodoListItem struct {
	Id          string `json:"id" bson:"_id,omitempty"`
	Item        string `json:"item" bson:"item,omitempty"`
	IsCompleted bool   `json:"isCompleted" bson:"isCompleted,omitempty"`
}
