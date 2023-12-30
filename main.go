package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the index page"))
}

func main() {
	log.Printf("Welcome to Dockerized Go Web Application\n")
	http.HandleFunc("/", indexHandler)
	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
