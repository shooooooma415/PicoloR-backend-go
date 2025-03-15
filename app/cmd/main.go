package main

import (
	"log"
	"net/http"

	v1 "picolor-backend/app/server/v1"
)

func main() {
	router := v1.NewRouter()

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}