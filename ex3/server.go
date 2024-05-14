package main

import (
	"log"
	"net/http"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	handler := newHandler(db)

	routes := http.NewServeMux()

	routes.HandleFunc("GET /get", handler.getHandler)

	log.Fatal(http.ListenAndServe("localhost:3020", routes))
}
