package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//Handler habitual
func holaMundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola mundo</h1>")
}
func holaLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Login</h1>")
}

//Handler avanzado
type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}

func main() {
	msg := mensaje{
		msg: "Hola mundo",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", holaMundo)
	mux.HandleFunc("/login", holaLogin)

	mux.Handle("/hola", msg)

	// http.ListenAndServe(":8080", mux)
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())

}
