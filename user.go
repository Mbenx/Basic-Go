package main

import (
	"fmt"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "All User Called")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Create User Called")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "update User Called")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Delete User Called")
}
