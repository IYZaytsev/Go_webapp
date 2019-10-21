package handlefunc

import (
	"net/http"
	"strconv"
	"strings"
)

//UpdateHandler Used to generate generate a commit
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	//grabs results from URL to createResult page and delivers results with switch
	URLSubString := r.URL.Path[len("/update/"):]
	URL := strings.Split(URLSubString, "/")

	databaseType, ID := URL[0], URL[1]

	switch databaseType {
	case "student":
		_, classList := APICallClass(w, r, "/classes/search", "GET", "//")
		_, studentList := APICallStudent(w, r, "/students/search", "GET", "//"+ID)

		selectString := "<h1> Editing " + studentList[0].Name + "</h1> Enroll In Class<select name = 'class_list'>"
		selectString += "<option value = none>none</option>"
		for i := range classList {
			selectString += "<option value=\"" + classList[i].Name + "\">" + classList[i].Name + "</option>"
		}
		selectString += "</select> <div>StudentName: <textarea name='studentName' rows='1' cols='15'>" + studentList[0].Name + "</textarea></div>"
		selectString += "<input type='hidden' id='ID' name='ID' value=" + strconv.Itoa(studentList[0].ID) + ">"
		data := Page{Title: "Edit Page", Content: selectString, Student: true, Type: "student"}
		TemplateInit(w, "update.html", data)

	case "teacher":
		_, classList := APICallClass(w, r, "/classes/search", "GET", "//")
		_, teacherList := APICallTeacher(w, r, "/teachers/search", "GET", "//"+ID)

		selectString := "<h1> Editing " + teacherList[0].Name + "</h1> Assign Class<select name = 'class_list'>"
		selectString += "<option value = none>none</option>"
		for i := range classList {
			selectString += "<option value=\"" + classList[i].Name + "\">" + classList[i].Name + "</option>"
		}
		selectString += "</select> <div>Teacher: <textarea name='teacherName' rows='1' cols='15'>" + teacherList[0].Name + "</textarea></div>"
		selectString += "<input type='hidden' id='ID' name='ID' value=" + strconv.Itoa(teacherList[0].ID) + ">"
		data := Page{Title: "Edit Page", Content: selectString, Teacher: true, Type: "teacher"}
		TemplateInit(w, "update.html", data)

	case "class":
		_, classList := APICallClass(w, r, "/classes/search", "GET", "//"+ID)

		selectString := "<h1> Editing " + classList[0].Name + "</h1>"
		selectString += "<div>ClassName: <textarea name='className' rows='1' cols='15'>" + classList[0].Name + "</textarea></div>"
		selectString += "<input type='hidden' id='ID' name='ID' value=" + strconv.Itoa(classList[0].ID) + ">"
		data := Page{Title: "Edit Page", Content: selectString, Class: true, Type: "class"}
		TemplateInit(w, "update.html", data)
	}
}

//CommitHandler gets passed along data to be changed
func CommitHandler(w http.ResponseWriter, r *http.Request) {
	urlSubString := r.URL.Path[len("/commitchanges/"):]

	switch urlSubString {
	case "student":
		r.ParseForm() // Required if you don't call r.FormValue()
		classChoice := strings.Join(r.Form["class_list"], " ")
		parameter := classChoice + "/" + r.FormValue("ID") + "/" + r.FormValue("studentName")

		APICallStudent(w, r, "/students/update", "PUT", parameter)
		http.Redirect(w, r, "/search/result/student", http.StatusSeeOther)
	case "teacher":
		r.ParseForm() // Required if you don't call r.FormValue()
		classChoice := strings.Join(r.Form["class_list"], " ")
		parameter := classChoice + "/" + r.FormValue("ID") + "/" + r.FormValue("teacherName")

		APICallTeacher(w, r, "/teachers/update", "PUT", parameter)
		http.Redirect(w, r, "/search/result/teacher", http.StatusSeeOther)

	case "class":

		parameter := r.FormValue("ID") + "/" + r.FormValue("className")

		APICallClass(w, r, "/classes/update", "PUT", parameter)
		http.Redirect(w, r, "/search/result/class", http.StatusSeeOther)
	}

}
