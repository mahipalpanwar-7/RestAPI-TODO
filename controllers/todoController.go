package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mahipalpanwar-7/RestAPI-TODO/models"
)

var todoList []models.Todo
var nextId = 1

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todoList)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid entry")
		return
	}

	if todo.Title == "" {
		json.NewEncoder(w).Encode("Title cannot be empty")
		return
	}
	todo.Id = nextId
	todo.Iscompleted = false
	nextId++

	todoList = append(todoList, todo)

	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode("invalid id")
	}

	for index, todo := range todoList {
		if todo.Id == id {
			todoList = append(todoList[:index], todoList[index+1:]...)

			json.NewEncoder(w).Encode("Todo deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("todo not found")
}

func CompleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Invalid Id")
		return
	}

	for index, todo := range todoList {
		if todo.Id == id {

			todoList[index].Iscompleted = true

			json.NewEncoder(w).Encode(todoList[index])

			return
		}
	}
	json.NewEncoder(w).Encode("Todo not found")
}
