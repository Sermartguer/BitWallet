package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"./common"
	"./users"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Creds struct {
	Status      string
	APIKey      string
	AccountType string
	Email       string
	AuthToken   string
	IsLoggedIn  bool
}

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"create_at"`
}

type jwtKey struct {
	jwt string
}

var noteStore = make(map[string]Note)

var id int

//GetNoteHandler - GET -	/ap/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var notes []Note

	db := common.DbConn()
	selDB, err := db.Query("SELECT * FROM test")
	if err != nil {
		panic(err.Error())
	}
	emp := Note{}
	for selDB.Next() {
		var id int
		var title, description string
		var create_at string
		err = selDB.Scan(&id, &title, &description, &create_at)
		if err != nil {
			panic(err.Error())
		}
		emp.Title = title
		emp.Description = description
		emp.CreatedAt, err = time.Parse(time.RFC3339, create_at)
		notes = append(notes, emp)
	}

	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//PostNoteHandler - POST -	/api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreatedAt = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note

	db := common.DbConn()

	insForm, err := db.Prepare("INSERT INTO test (title, description, create_at) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(note.Title, note.Description, note.CreatedAt)
	log.Println("INSERT: Title: " + note.Title + " | Description: " + note.Description + " | CreateAt: " + note.CreatedAt.String())

	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//PutNoteHandler - PUT -	/ap/notes
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	var noteUpdate Note
	err := json.NewDecoder(r.Body).Decode(&noteUpdate)
	if err != nil {
		panic(err)
	}
	if note, ok := noteStore[k]; ok {
		noteUpdate.CreatedAt = note.CreatedAt
		delete(noteStore, k)
		noteStore[k] = noteUpdate
	} else {
		log.Printf("No encontamos la nota")
	}
	w.WriteHeader(http.StatusNoContent)
}

//DeleteUsersHandler - DELETE -	/ap/notes
func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]

	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		log.Printf("No encontamos la nota")
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteUsersHandler).Methods("DELETE")
	r.HandleFunc("/api/login", users.Login).Methods("POST")
	r.HandleFunc("/api/isLoged", users.Islogged).Methods("POST")
	server := &http.Server{
		Addr:           ":8080",
		Handler:        cors.Default().Handler(r),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening http://localhost:8080 ...")
	log.Fatal(server.ListenAndServe())
}
