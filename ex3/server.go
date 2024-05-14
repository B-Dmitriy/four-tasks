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

	usersHandler := newUsersHandler(db)

	routes := http.NewServeMux()

	routes.HandleFunc("GET /users", usersHandler.getUsersList)

	log.Fatal(http.ListenAndServe("localhost:3020", routes))
}
