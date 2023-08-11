package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

var db *gorm.DB
var err error

type Student struct {
	gorm.Model
	Name string `json:"name"`
	Roll string `json:"roll"`
}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Printf(err.Error())
		panic("Failed to connect to the server")
	}
	defer db.Close()

	db.AutoMigrate(&Student{})

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to Database")
	}
	defer db.Close()

	params := mux.Vars(r)
	roll := params["roll"]

	var student Student
	db.Where("roll = ?", roll).Find(&student)
	db.Delete(&student)

	json.NewEncoder(w).Encode(student)

}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to Database")
	}
	defer db.Close()

	var newStudent Student
	json.NewDecoder(r.Body).Decode(&newStudent)

	db.Create(&newStudent)

	json.NewEncoder(w).Encode(newStudent)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to Database")
	}
	defer db.Close()

	params := mux.Vars(r)
	roll := params["roll"]

	var student Student
	db.Where("roll = ?", roll).Find(&student)

	json.NewDecoder(r.Body).Decode(&student)
	student.Roll = roll
	db.Save(&student)

	json.NewEncoder(w).Encode(student)
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to Database")
	}
	defer db.Close()

	var students []Student
	db.Find(&students)
	json.NewEncoder(w).Encode(students)
}

func GetStudentByRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to Database")
	}
	defer db.Close()

	params := mux.Vars(r)
	roll := params["roll"]

	var student Student
	db.Where("roll = ?", roll).Find(&student)

	json.NewEncoder(w).Encode(student)
}
