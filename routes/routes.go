package routes

import (
	"github.com/mahipalpanwar-7/RestAPI-TODO/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("todo", controllers.AddTodo).Methods("POST")
	router.HandleFunc("todos/{id}", controllers.CompleteTodo).Methods("PUT")
	router.HandleFunc("todos/{id}", controllers.DeleteTodo).Methods("DELETE")
	return router
}
