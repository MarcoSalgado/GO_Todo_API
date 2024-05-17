package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marco/todoapi/handlers"
)

func main() {
	app := handlers.App{
		NextID: 1,
		Todos:  make([]handlers.Todo, 0),
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r, &app)
	}).Methods("GET")

	router.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTodoHandler(w, r, &app)
	}).Methods("POST")

	router.HandleFunc("/todos/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTodoHandler(w, r, &app)
	}).Methods("GET")

	router.HandleFunc("/todos/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateTodoHandler(w, r, &app)
	}).Methods("PUT")

	router.HandleFunc("/todos/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTodoHandler(w, r, &app)
	}).Methods("DELETE")

	router.HandleFunc("/todos/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllTodosHandler(w, r, &app)
	}).Methods("GET")

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
