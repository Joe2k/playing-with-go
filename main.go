package main

import (
	"log"
	"net/http"

	"github.com/Joe2k/playing-with-go/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/notification", handlers.GetNotifications).Methods("GET")
	r.HandleFunc("/api/notification", handlers.CreateNotification).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
