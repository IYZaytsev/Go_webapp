package handlefunc

import (
	"net/http"
	"strconv"
	"strings"
)

//UpdateHandler Used to generate generate a commit
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	//grabs results from URL to createResult page and delivers results with switch
	urlSubString := r.URL.Path[len("/update/"):]
	URl := strings.Split(urlSubString, "/")

	databaseType, ID := URl[0], URl[1]

	switch databaseType {
	case "student":
		_, classList := APICallClass(w, r, "/classes/search", "GET", "//")
		_, studentList := APICallStudent(w, r, "/students/search", "GET", "//"+ID)

		selectString := "<h1> Editing " + studentList[0].Name + "</h1> Enroll In Class<select name = 'class_list'>"
		selectString += "<option value = none>none</option>"
		for i := range classList {
			selectString += "<option value=" + classList[i].Name + ">" + classList[i].Name + "</option>"
		}
		selectString += "</select> <div>StudentName: <textarea name='studentName' rows='1' cols='15'>" + studentList[0].Name + "</textarea></div>"
		selectString += "<input type='hidden' id='ID' name='ID' value=" + strconv.Itoa(studentList[0].ID) + ">"
		data := Page{Title: "Edit Page", Content: selectString, Student: true, Type: "student"}
		TemplateInit(w, "update.html", data)

	case "teacher":

	case "class":
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
		//http.Redirect(w, r, "/search/result/student", http.StatusSeeOther)
	}

}
