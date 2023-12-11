package routes

import "net/http"

func Router(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch r.URL.Path {
	case "/":
	//do something
	case "/tasks":
		//Get all the tasks from db
		if method == http.MethodGet {
			return
		}
		//Creating a task
		if method == http.MethodPost {
			return
		}
		http.Error(w, "Bad request", http.StatusBadRequest)
	case "/tasks/id":
		//delete a task
		if method == http.MethodDelete {
			return
		}
		//update a task
		if method == http.MethodPut {
			return
		}
		//get/read a specific task
		if method == http.MethodGet {
			return
		}
		http.Error(w, "Bad request", http.StatusBadRequest)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}
