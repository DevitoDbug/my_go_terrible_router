package controllers

import (
	"encoding/json"
	"log"
	"my_go_terrible_router/pkg/models"
	"my_go_terrible_router/pkg/utility"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	task := &models.Task{}
	// Read the request body
	err := utility.ReadBody(r, task)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//adding the task to db
	err = task.CreateTask()
	if err != nil {
		log.Printf("Error in adding task to db : %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Encode the task directly to the response writer
	enc := json.NewEncoder(w)
	err = enc.Encode(task)
	if err != nil {
		log.Printf("Error encoding task: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetAllTasks()
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	err = enc.Encode(tasks)
	if err != nil {
		log.Printf("Error encoding task: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {

}
func DeleteTask(w http.ResponseWriter, r *http.Request) {

}
func UpdateTask(w http.ResponseWriter, r *http.Request) {

}
