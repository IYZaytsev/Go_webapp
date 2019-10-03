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
var tmpls = template.Must(template.ParseFiles("tmpl/index.html", "tmpl/select.html", "tmpl/create.html", "tmpl/view.html",
	"tmpl/search.html", "tmpl/searchResults.html"))

//Used to create unique ids for items
var globalID int

//used to change title and header of a HTML template

type page struct {
	Title   string
	Header  string
	Action  string
	Teacher bool
	Student bool
	Class   bool
	Type    string
	Content string
	ID      int
}

/*
func processResult(r results) string {
	name := "Name: " + r.name + " ID: "
	id := r.id
	className := " Classes: "
	for i := 0; i < len(r.classes); i++ {
		className = className + r.classes[i].Name
	}
	returnString := fmt.Sprintf("%s%d%s", name, id, className)
	return returnString

}
*/
func generateUniqueID() int {
	globalID++
	return globalID
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
	urlSubString := r.URL.Path[len("/select/"):]

	switch urlSubString {
	case "create":
		data := page{Title: "School Database", Header: "What type would you like to create?", Action: "/" + urlSubString + "/"}
		templateInit(w, "select.html", data)

	case "search":
		data := page{Title: "School Database", Header: "What type would you like to search?", Action: "/" + urlSubString + "/"}
		templateInit(w, "select.html", data)

	}

}
func searchHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() // Required if you don't call r.FormValue()
	value := strings.Join(r.Form["selected_value"], " ")

	switch value {
	case "student":
		data := page{Title: "search Page", Content: "student search page", Student: true, Type: "student"}
		templateInit(w, "search.html", data)
	case "teacher":
		data := page{Title: "search Page", Content: "teacher search page", Teacher: true, Type: "teacher"}
		templateInit(w, "search.html", data)
	case "class":
		data := page{Title: "search Page", Content: "class search page", Class: true, Type: "class"}
		templateInit(w, "search.html", data)
	}

}
func createHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Required if you don't call r.FormValue()
	value := strings.Join(r.Form["selected_value"], " ")
	log.Println(r.URL.Path)
	data := page{Title: "School Database", Header: "Editing " + value, Type: value}

	switch value {
	case "student":
		data.Student = true

	case "teacher":
		data.Teacher = true

	case "class":
		data.Class = true

	}

	templateInit(w, "create.html", data)

}

//main page of the web app
func index(w http.ResponseWriter, r *http.Request) {
	data := page{Title: "School Database", Header: "Welcome, please select an option"}
	templateInit(w, "index.html", data)

}

/*
func visualizeHandler(w http.ResponseWriter, r *http.Request) {

}
*/
func viewHandler(w http.ResponseWriter, r *http.Request) {
	typeOfObject := r.URL.Path[len("/view/"):]

	switch typeOfObject {

	case "student":

		tempID := generateUniqueID()
		dataBaseTypes.AddStudent(tempID, r.FormValue("studentName"), &dataBaseStudent)
		data := page{Title: "View Page", Content: "Student: " + r.FormValue("studentName"), ID: tempID}
		templateInit(w, "view.html", data)

	case "teacher":

		tempID := generateUniqueID()
		var classSlice []dataBaseTypes.Class
		dataBaseTypes.AddTeacher(tempID, classSlice, r.FormValue("teacherName"), &dataBaseTeacher)
		data := page{Title: "View Page", Content: "Teacher: " + r.FormValue("teacherName"), ID: tempID}
		templateInit(w, "view.html", data)

	case "class":

		tempID := generateUniqueID()
		var studentSlice []dataBaseTypes.Student
		dataBaseTypes.AddClass(tempID, r.FormValue("className"), studentSlice, -1, &dataBaseClass)
		data := page{Title: "View Page", Content: "className: " + r.FormValue("className"), ID: tempID}
		templateInit(w, "view.html", data)

	} //end of switch statement
}

/*func resultsHandler(w http.ResponseWriter, r *http.Request) {
	typeOfObject := r.URL.Path[len("/searchResults/"):]
	switch typeOfObject {
	case "student":
		result := searchDatabase(r.FormValue("studentName"))
		resultString := processResult(result)
		data := page{Title: "results Page", Content: resultString}
		log.Println(result)
		result.Student = true
		templateInit(w, "searchResults.html", data)

	}

}
*/
func main() {
	/*	dataBaseTypes.AddStudent(69, "bob saget", &dataBaseStudent)
		var tempSliceStudent []dataBaseTypes.Student
		tempSliceStudent = append(tempSliceStudent, dataBaseStudent[0])
		dataBaseTypes.AddClass(32, "history", tempSliceStudent, 23, &dataBaseClass)
		fmt.Println(len(dataBaseClass))
		fmt.Println(len(dataBaseStudent))
	*/
	http.HandleFunc("/", index)
	http.HandleFunc("/select/", selectHandler)
	http.HandleFunc("/create/", createHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/search/", searchHandler)
	//http.HandleFunc("/searchResults/", resultsHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("listening")
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
