package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

type simpleHandler struct{}

func (sh *simpleHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	res := &Response{
		Message: "Hello, World!",
		Time:    time.Now().Format(time.RFC3339),
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

type healthHandler struct{}

func (hh *healthHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	res := &struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /api/hello", &simpleHandler{})
	mux.Handle("GET /health", &healthHandler{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
