package main

import (
	"fmt"
	"net/http"

	"github.com/mahipalpanwar-7/RestAPI-TODO/routes"
)

func main() {
	router := routes.Router()
	fmt.Println("listening to the server")
	http.ListenAndServe(":8080", router)
}
