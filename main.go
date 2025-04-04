package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GOAT-stack-todo-app/routes"
)

func main() {

	fmt.Println("GOAT todo-app")

	// init the server.
	server := http.NewServeMux()
	PORT := ":5000"

	// static landing page
	fs := http.FileServer(http.Dir("public"))
	server.Handle("/public/", http.StripPrefix("/public/", fs))
	// api endpoints
	server.HandleFunc("GET /{$}", routes.HomePage)
	server.HandleFunc("POST /add-task", routes.AddTask)
	server.HandleFunc("GET /tasks", routes.GetAllTasks)
	server.HandleFunc("GET /finish/{id}", routes.FinishByID)
	server.HandleFunc("GET /delete/{id}", routes.DeleteByID)

	log.Println("running on >> http://127.0.0.1:5000")
	log.Fatal(http.ListenAndServe(PORT, server))
}
