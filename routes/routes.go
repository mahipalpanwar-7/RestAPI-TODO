package routes

import "github.com/gorilla/mux"

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("todos", GetTodos).Methods("GET")
	router.HandleFunc("todo", AddTodo).Methods("POST")
	router.HandleFunc("todos/{id}", UpdateTodo).Methods("PUT")
	router.HandleFunc("todos/{id}", DeleteTodo).Methods("DELETE")
	return router
}
