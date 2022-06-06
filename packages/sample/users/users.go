package main

import (
	"net/http"

	"encoding/json"
	"github.com/gorilla/mux"
)

func Main() {

	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/user/{id}/score/{score}", postUser).Methods("POST")
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

	r.HandleFunc("/testing", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("function is working"))
	})

	// http.ListenAndServe(":8080", r)

}

type User struct {
	Id    string
	Score int
}

// /user/{id}
func getUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userid := params["id"]

	// fake user
	testUser := User{
		Id:    userid,
		Score: 10,
	}

	userJSON, err := json.Marshal(testUser)
	if err != nil {
		panic(err)
	}

	w.Write(userJSON)
}

// /user/{id}/score/{score}
func postUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userid := params["id"]
	score := params["score"]

	// this does not update any data source. This is scaffolding
	w.Write([]byte("posted user " + userid + " with score " + score))
}

// /user/{id}
func deleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userid := params["id"]

	// this does not update any data source. This is scaffolding
	w.Write([]byte("deleted " + userid))
}
