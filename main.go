package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")

}

func handleRequest() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HelloWorld).Methods("GET")
	// Users endpoint

	router.HandleFunc("/users", AllUsers).Methods("GET")
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	fmt.Println("Go Run On Port 8081")

	handleRequest()
}
