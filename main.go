package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type UsersContainer struct {
	Users []User `json:"users"`
}

type User struct {
	ID   int
	Name string
}

func Get(w http.ResponseWriter, r *http.Request) {
	var users UsersContainer
	jsonFile, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(byteValue, &users); err != nil {
		// TODO: handle error
	}
	if err != nil {
		// TODO: handle error
	}

	result := users
	if err := json.NewEncoder(w).Encode(result); err != nil {
		// TODO: handle error
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", Get).
		Methods("GET")
	fmt.Println("starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
