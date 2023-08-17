package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Api B is working!"))
}

func getGreetings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Greetings from api B!"))
}

func main() {
	r := mux.NewRouter()

	s := r.PathPrefix("/api-b").Subrouter()
	s.Use(logMiddleware)
	s.HandleFunc("/health_check", HealthCheckHandler)
	s.HandleFunc("/greetings", getGreetings)
	http.ListenAndServe(":8080", r)
}
