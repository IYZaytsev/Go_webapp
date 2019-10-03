package handlefunc

import (
	"net/http"
	"strings"

	"../databasetypes"
)

var globalID int

//generateUniqueID is used to create unique IDS for dataBase entries
func generateUniqueID() int {
	globalID++
	return globalID
}

//SearchHandler Used to generate search page based on the results of the select page
func SearchHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() // Required if you don't call r.FormValue()
	value := strings.Join(r.Form["selected_value"], " ")

	switch value {
	case "student":
		data := Page{Title: "search Page", Content: "student search page", Student: true, Type: "student"}
		TemplateInit(w, "search.html", data)
	case "teacher":
		data := Page{Title: "search Page", Content: "teacher search page", Teacher: true, Type: "teacher"}
		TemplateInit(w, "search.html", data)
	case "class":
		data := Page{Title: "search Page", Content: "class search page", Class: true, Type: "class"}
		TemplateInit(w, "search.html", data)
	}

}

//CreateHandler used to create createPage based on the value of the select page
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Required if you don't call r.FormValue()
	value := strings.Join(r.Form["selected_value"], " ")
	data := Page{Title: "School Database", Header: "Editing " + value, Type: value}

	switch value {
	case "student":
		data.Student = true

	case "teacher":
		data.Teacher = true

	case "class":
		data.Class = true

	}

	TemplateInit(w, "create.html", data)

}

//ViewHandler used to create html for newly created objects based on the createPage,
//also creates objects in the database
func ViewHandler(w http.ResponseWriter, r *http.Request) {

	typeOfObject := r.URL.Path[len("/view/"):]

	switch typeOfObject {

	case "student":

		tempID := generateUniqueID()
		databasetypes.AddStudent(tempID, r.FormValue("studentName"), &dataBaseStudent)
		data := Page{Title: "View Page", Content: "Student: " + r.FormValue("studentName"), ID: tempID}
		TemplateInit(w, "view.html", data)

	case "teacher":

		tempID := generateUniqueID()
		var classSlice []databasetypes.Class
		databasetypes.AddTeacher(tempID, classSlice, r.FormValue("teacherName"), &dataBaseTeacher)
		data := Page{Title: "View Page", Content: "Teacher: " + r.FormValue("teacherName"), ID: tempID}
		TemplateInit(w, "view.html", data)

	case "class":

		tempID := generateUniqueID()
		var studentSlice []databasetypes.Student
		databasetypes.AddClass(tempID, r.FormValue("className"), studentSlice, -1, &dataBaseClass)
		data := Page{Title: "View Page", Content: "className: " + r.FormValue("className"), ID: tempID}
		TemplateInit(w, "view.html", data)

	} //end of switch statement
}
