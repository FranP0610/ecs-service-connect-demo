package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Api B is working!"))
}

func getGreetings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Greetings from api B!"))
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/api-b").Subrouter()
	//r.PathPrefix("/api-a/")
	s.HandleFunc("/health_check", HealthCheckHandler)
	s.HandleFunc("/greetings", getGreetings)
	http.ListenAndServe(":8080", r)
}
