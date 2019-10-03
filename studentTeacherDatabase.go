package main

import (
	"log"
	"net/http"

	"./databasetypes"

	"./handlefunc"
)

//DataBaseStudent used to store information on student
var DataBaseStudent []databasetypes.Student

//DataBaseClass used to store information on student
var DataBaseClass []databasetypes.Class

//DataBaseTeacher used to store information on student
var DataBaseTeacher []databasetypes.Teacher

func main() {

	handlefunc.SetHandlers(&DataBaseStudent, &DataBaseTeacher, &DataBaseClass)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("listening")
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
