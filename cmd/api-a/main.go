package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

// const urlApiB = "http://api-b:8080/api-b/greetings"
const urlApiB = "http://localhost:8080/api-b/greetings"

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Api A is working!"))
}

func getGreetings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Greetings from api A!"))
}

func callApiB(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(urlApiB)
	if err != nil {
		panic(err)
	}
	// Close the response body
	defer resp.Body.Close()
	// Read all the contents of the Reader - Buffer to read data
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	w.Write(body)
}
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
func main() {
	r := mux.NewRouter()
	//s.Use(logMiddleware)
	s := r.PathPrefix("/api-a").Subrouter()
	s.Use(logMiddleware)
	s.HandleFunc("/health_check", healthCheck)
	s.HandleFunc("/greetings", getGreetings)
	s.HandleFunc("/api-b", callApiB)
	http.ListenAndServe(":8000", s)
}
