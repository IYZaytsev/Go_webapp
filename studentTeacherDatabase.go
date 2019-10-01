package main

import (
	"html/template"
	"log"
	"net/http"

	"./dataBaseTypes"
)

//Creating slices of stu
var dataBaseStudent []dataBaseTypes.Student
var dataBaseClass []dataBaseTypes.Class
var dataBaseTeacher []dataBaseTypes.Teacher
var tmpls = template.Must(template.ParseFiles("tmpl/index.html", "tmpl/create.html"))

type Page struct {
	Title  string
	Header string
}

func createHandler(w http.ResponseWriter, r *http.Request) {

	data := Page{Title: "School Database", Header: "What type would you like to create?"}
	if err := tmpls.ExecuteTemplate(w, "create.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	data := Page{Title: "School Database", Header: "Welcome, please select an option"}
	if err := tmpls.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create/", createHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("listening")
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
