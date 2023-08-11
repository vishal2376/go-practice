package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "echo:Vishal@2376@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Successfully Connected to My SQL Database")

	insert, err := db.Query("INSERT INTO Users VALUES ('Go Lang');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successfully Inserted to My SQL Database")
}
