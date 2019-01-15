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
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	fmt.Println("Go Run On Port 8081")

	handleRequest()
}
