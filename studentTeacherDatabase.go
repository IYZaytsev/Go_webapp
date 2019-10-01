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
	Content string
	ID      int
}

//function to load template
func templateInit(w http.ResponseWriter, templateFile string, templateData page) {
	if err := tmpls.ExecuteTemplate(w, templateFile, templateData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//handler used for the first create screen
func selectHandler(w http.ResponseWriter, r *http.Request) {

	data := page{Title: "School Database", Header: "What type would you like to create?"}
	templateInit(w, "select.html", data)

}

func createHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Required if you don't call r.FormValue()
	value := strings.Join(r.Form["selected_value"], " ")
	data := page{Title: "School Database", Header: "Editing " + value, Type: value}

	switch value {
	case "student":
		data.Student = true

	case "teacher":
		data.Student = true

	case "class":
		data.Student = true

	}

	templateInit(w, "create.html", data)

}

//main page of the web app
func index(w http.ResponseWriter, r *http.Request) {
	data := page{Title: "School Database", Header: "Welcome, please select an option"}
	templateInit(w, "index.html", data)

}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("in view handler")
	typeOfObject := r.URL.Path[len("/view/"):]

	switch typeOfObject {

	case "student":
		log.Println("Inside student logic")
		dataBaseTypes.AddStudent(1, r.FormValue("studentName"), &dataBaseStudent)
		data := page{Title: "View Page", Content: "Student: " + r.FormValue("studentName"), ID: 1}
		log.Println("creating viewpage struct")
		templateInit(w, "view.html", data)

	case "teacher:":

	case "class":

	} //end of switch statement
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
