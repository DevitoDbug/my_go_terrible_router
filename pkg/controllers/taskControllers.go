package controllers

import (
	"encoding/json"
	"log"
	"my_go_terrible_router/pkg/models"
	"net/http"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetAllTasks()
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(tasks)
	if err != nil {
		log.Printf("Error in marshalling tasks: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(res))
	if err != nil {
		log.Printf("Error in response writting: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {

}

func CreateTask(w http.ResponseWriter, r *http.Request) {

}
func DeleteTask(w http.ResponseWriter, r *http.Request) {

}
func UpdateTask(w http.ResponseWriter, r *http.Request) {

}
