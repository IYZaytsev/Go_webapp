package handlefunc

import (
	"net/http"

	"../databasetypes"
)

var dataBaseStudent []databasetypes.Student
var dataBaseClass []databasetypes.Class
var dataBaseTeacher []databasetypes.Teacher

//SetHandlers cuts down on the amount of code in main by setting
// each url to a handlerfunction
func SetHandlers() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/select/", SelectHandler)
	http.HandleFunc("/create/", CreateHandler)
	http.HandleFunc("/view/", ViewHandler)
	http.HandleFunc("/search/", SearchHandler)
	http.HandleFunc("/update/", UpdateHandler)
	http.HandleFunc("/commitchanges/", CommitHandler)
	http.HandleFunc("/delete/", DeleteHandler)

}
