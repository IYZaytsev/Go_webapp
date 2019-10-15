package handlefunc

import (
	"net/http"
	"strconv"
	"strings"
)

//DeleteHandler Used to generate generate a commit
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	//grabs results from URL to createResult page and delivers results with switch
	URLSubString := r.URL.Path[len("/delete/"):]
	URL := strings.Split(URLSubString, "/")

	databaseType, ID := URL[0], URL[1]

	switch databaseType {
	case "student":
		APICallStudent(w, r, "/students/delete", "DELETE", ID)
		//http.Redirect(w, r, "/search/result/student", http.StatusSeeOther)
	case "teacher":
		_, classList := APICallClass(w, r, "/classes/search", "GET", "//")
		_, teacherList := APICallTeacher(w, r, "/teachers/search", "GET", "//"+ID)

		selectString := "<h1> Editing " + teacherList[0].Name + "</h1> Assign Class<select name = 'class_list'>"
		selectString += "<option value = none>none</option>"
		for i := range classList {
			selectString += "<option value=" + classList[i].Name + ">" + classList[i].Name + "</option>"
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
