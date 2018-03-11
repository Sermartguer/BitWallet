package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde GET")
}
func POSTUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde POST")
}
func PUTUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde PUT")
}
func DELETEUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde Delete")
}
func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users", POSTUsers).Methods("POST")
	r.HandleFunc("/api/users", PUTUsers).Methods("PUT")
	r.HandleFunc("/api/users", DELETEUsers).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())

}
