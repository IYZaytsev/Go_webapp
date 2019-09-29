package main

import (
	"log"
	"net/http"

	"./dataBaseTypes"
)

//Creating slices of stu
var dataBaseStudent []dataBaseTypes.Student
var dataBaseClass []dataBaseTypes.Class
var dataBaseTeacher []dataBaseTypes.Teacher

func main() {
	fs := http.FileServer(http.Dir("statics"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
