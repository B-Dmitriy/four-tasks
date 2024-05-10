package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	routes := http.NewServeMux()

	routes.HandleFunc("POST /redirection", redirectionHandler)
	routes.HandleFunc("GET /get", getHandler)

	log.Fatal(http.ListenAndServe("localhost:3010", routes))
}

// redirectionHandler - получает json массив вида {1:"one",3:"two",2:"four"}
// Возвращает ответ со статусом 200 (StatusOk) c этим json в отсортированном и обычном виде
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

// getHandler ждёт в параметрах token, если находит - возвращает ответ со статусом 200,
// в противном случае возвращает статус 400 (BadRequest)
func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Method unsupported"))
		return
	}

	if r.URL.RawQuery == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryParams, _ := url.ParseQuery(r.URL.RawQuery)

	if len(queryParams["token"]) == 0 || queryParams["token"][0] == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("need token in query params\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
