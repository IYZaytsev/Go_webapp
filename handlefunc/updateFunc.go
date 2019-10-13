package handlefunc

import (
	"net/http"
	"strings"
)

//UpdateHandler Used to generate search page based on the results of the select page
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	//grabs results from URL to createResult page and delivers results with switch
	urlSubString := r.URL.Path[len("/update/"):]
	URl := strings.Split(urlSubString, "/")

	databaseType, ID := URl[0], URl[1]

	switch databaseType {
	case "student":
		_, classList := APICallClass(w, r, "/classes/search", "GET", "//")
		_, studentList := APICallStudent(w, r, "/students/search", "GET", "//"+ID)

		/*	<select name = "selected_value">
				<option  value="student">student</option>
				<option value="teacher">teacher</option>
				<option value="class">class</option>
			</select>
		*/
		selectString := "<h1> Editing " + studentList[0].Name + "</h1>" + "<h1>Enroll in Class: <select name = 'class_list'>"
		for i := range classList {
			selectString += "<option value=" + classList[i].Name + ">" + classList[i].Name + "</option>"
		}
		selectString += "</select>"
		data := Page{Title: "Edit Page", Content: selectString, Student: true}
		TemplateInit(w, "update.html", data)

	case "teacher":

	case "class":
	}
}
