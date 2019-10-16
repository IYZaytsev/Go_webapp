package handlefunc

import (
	"net/http"
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
		http.Redirect(w, r, "/search/result/student", http.StatusSeeOther)
	case "teacher":
		APICallStudent(w, r, "/teachers/delete", "DELETE", ID)
		http.Redirect(w, r, "/search/result/teacher", http.StatusSeeOther)
	case "class":
		APICallStudent(w, r, "/classes/delete", "DELETE", ID)
		http.Redirect(w, r, "/search/result/class", http.StatusSeeOther)
	}
}
