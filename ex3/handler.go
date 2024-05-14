package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

const UsersLimit = 10

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type GetHandler struct {
	db *sql.DB
}

func newHandler(db *sql.DB) *GetHandler {
	return &GetHandler{
		db: db,
	}
}

func (h *GetHandler) getHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("SELECT * FROM users LIMIT ?;", UsersLimit)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	users := make([]User, 0, UsersLimit)
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		users = append(users, u)
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(usersJSON)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}
