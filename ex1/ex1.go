package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
/redirection POST - Выполняет запрос на указанный адрес и получает json массив вида {1:"one",3:"two",2:"four"}
Вернуть ответ со статусом 200 (StatusOk) c этим json в отсортированном и обычном виде
/get GET если в параметрах передается token вернуть ответ со статусом 200, в противном случае вернуть статус 400 (BadRequest)
*/
func main() {
	routes := http.NewServeMux()

	routes.HandleFunc("POST /redirection", redirectionHandler)
	routes.HandleFunc("GET /get", getHandler)

	log.Fatal(http.ListenAndServe("localhost:3010", routes))
}

func redirectionHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sortedJSONString, err := sortBody(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(createRedirectionResponse(body, sortedJSONString))
	if err != nil {
		w.WriteHeader(500)
		return
	}
}

func createRedirectionResponse(origin, sorted []byte) []byte {
	return []byte(fmt.Sprintf("{ \"origin\": %s, \"sort\": %s }", origin, sorted))
}

func sortBody(body []byte) ([]byte, error) {
	sortedMap := make(map[string]string)
	err := json.Unmarshal(body, &sortedMap)
	if err != nil {
		return nil, err
	}

	sortedJSONString, err := json.Marshal(sortedMap)
	if err != nil {
		return nil, err
	}

	return sortedJSONString, nil
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Method unsupported"))
		return
	}
}
