package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"./dataBaseTypes"
)

//Global int variable to create unqiue ids

//Creating slices of different types
var dataBaseStudent []dataBaseTypes.Student
var dataBaseClass []dataBaseTypes.Class
var dataBaseTeacher []dataBaseTypes.Teacher
var tmpls = template.Must(template.ParseFiles("tmpl/index.html", "tmpl/select.html", "tmpl/create.html", "tmpl/view.html"))

//used to change title and header of a HTML template
type page struct {
	Title   string
	Header  string
	Teacher bool
	Student bool
	Class   bool
	Type    string
}

//handler used for the first create screen

func selectHandler(w http.ResponseWriter, r *http.Request) {

	data := page{Title: "School Database", Header: "What type would you like to create?"}
	if err := tmpls.ExecuteTemplate(w, "select.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Required if you don't call r.FormValue()
	value := strings.Join(r.Form["selected_value"], " ")
	data := page{Title: "School Database", Header: "Editing " + value, Type: value}

	if value == "student" {
		data.Student = true

	}
	if value == "teacher" {
		data.Teacher = true

	}
	if value == "class" {
		data.Class = true

	}
	if err := tmpls.ExecuteTemplate(w, "create.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//main page of the web app
func index(w http.ResponseWriter, r *http.Request) {
	data := page{Title: "School Database", Header: "Welcome, please select an option"}
	if err := tmpls.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	typeOfObject := r.URL.Path[len("/view/"):]
	if typeOfObject == "student" {
		log.Println(r.FormValue("studentName"))
		dataBaseTypes.AddStudent(1, r.FormValue("studentName"), &dataBaseStudent)
	}

}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/select/", selectHandler)
	http.HandleFunc("/create/", createHandler)
	http.HandleFunc("/view/", viewHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("listening")
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
