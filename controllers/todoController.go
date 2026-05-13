package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mahipalpanwar-7/RestAPI-TODO/models"
)

var todoList []models.Todo
var nextId = 1

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todoList)
}
