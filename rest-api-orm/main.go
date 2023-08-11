package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/students", GetAllStudents).Methods("GET")
	myRouter.HandleFunc("/students", CreateStudent).Methods("POST")
	myRouter.HandleFunc("/students/{roll}", GetStudentByRoll).Methods("GET")
	myRouter.HandleFunc("/students/{roll}", DeleteStudent).Methods("DELETE")
	myRouter.HandleFunc("/students/{roll}", UpdateStudent).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To Student API Home Page")
}

func main() {
	println("GO ORM Practice - Student API")

	InitialMigration()

	handleRequests()
}
