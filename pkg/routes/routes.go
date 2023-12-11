package routes

import (
	"my_go_terrible_router/pkg/controllers"
	"net/http"
)

func Router(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch r.URL.Path {
	case "/":
	//do something
	case "/tasks":
		//Get all the tasks from db
		if method == http.MethodGet {
			controllers.GetAllTasks(w, r)
			return
		}
		//Creating a task
		if method == http.MethodPost {
			controllers.CreateTask(w, r)
			return
		}
		http.Error(w, "Bad request", http.StatusBadRequest)
	case "/tasks/id":
		//delete a task
		if method == http.MethodDelete {
			controllers.DeleteTask(w, r)
			return
		}
		//update a task
		if method == http.MethodPut {
			controllers.UpdateTask(w, r)
			return
		}
		//get/read a specific task
		if method == http.MethodGet {
			controllers.GetTask(w, r)
			return
		}
		http.Error(w, "Bad request", http.StatusBadRequest)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}
