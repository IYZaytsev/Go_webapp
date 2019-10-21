package handlefunc

import (
	"net/http"
	"strings"
)

var globalID int

//generateUniqueID is used to create unique IDS for dataBase entries
func generateUniqueID() int {
	globalID++
	return globalID
}

//CreateHandler used to create createPage based on the value of the select page
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Required if you don't call r.FormValue()
	value := strings.Join(r.Form["selected_value"], " ")
	data := Page{Title: "School Database", Header: "Creating " + value, Type: value}

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
		studentName := r.FormValue("studentName")
		studentStruct, _ := APICallStudent(w, r, "/students/create", "POST", studentName)
		data := Page{Title: "View Page", Content: "Student: " + r.FormValue("studentName"), ID: studentStruct.ID}
		TemplateInit(w, "view.html", data)

	case "teacher":
		teacherName := r.FormValue("teacherName")
		teacherStruct, _ := APICallTeacher(w, r, "/teachers/create", "POST", teacherName)
		data := Page{Title: "View Page", Content: "Teacher: " + r.FormValue("teacherName"), ID: teacherStruct.ID}
		TemplateInit(w, "view.html", data)

	case "class":

		className := r.FormValue("className")
		classStruct, _ := APICallClass(w, r, "/classes/create", "POST", className)
		data := Page{Title: "View Page", Content: "className: " + r.FormValue("className"), ID: classStruct.ID}
		TemplateInit(w, "view.html", data)

	} //end of switch statement
}
