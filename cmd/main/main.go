package main

import (
	"fmt"
	"log"
	"my_go_terrible_router/pkg/routes"
	"net/http"
)

func main() {
	port := ":8080"
	var router http.HandlerFunc
	router = routes.Router

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Printf("listen and server error: %v", err)
		return
	}
	fmt.Println("Starting sever at port ", port, "....")
}
