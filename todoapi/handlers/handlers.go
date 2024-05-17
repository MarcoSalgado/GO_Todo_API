package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// App holds shared data for handlers
type App struct {
	NextID int
	Todos  []Todo
}

// Todo represents a task with a title and content
type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request, app *App) {
	fmt.Fprintf(w, "Welcome to the TODO API!")
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request, app *App) {
	var newTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTodo.ID = app.NextID
	app.NextID++
	app.Todos = append(app.Todos, newTodo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func GetTodoHandler(w http.ResponseWriter, r *http.Request, app *App) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, todo := range app.Todos {
		if todo.ID == id {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	http.Error(w, "TODO not found", http.StatusNotFound)
}

func GetAllTodosHandler(w http.ResponseWriter, r *http.Request, app *App) {
	type TodosResponse struct {
		TotalCount int    `json:"total_count"`
		Todos      []Todo `json:"todos"`
	}

	response := TodosResponse{
		TotalCount: len(app.Todos),
		Todos:      app.Todos,
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request, app *App) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, todo := range app.Todos {
		if todo.ID == id {
			app.Todos[i].Title = updatedTodo.Title
			app.Todos[i].Content = updatedTodo.Content
			json.NewEncoder(w).Encode(app.Todos[i])
			return
		}
	}
	http.Error(w, "TODO not found", http.StatusNotFound)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request, app *App) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, todo := range app.Todos {
		if todo.ID == id {
			app.Todos = append(app.Todos[:i], app.Todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "TODO not found", http.StatusNotFound)
}
