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
//It also is used to pass the slices
func SetHandlers(dbStudent *[]databasetypes.Student, dbTeacher *[]databasetypes.Teacher, dbClass *[]databasetypes.Class) {
	dataBaseStudent = *dbStudent
	dataBaseTeacher = *dbTeacher
	dataBaseClass = *dbClass

	http.HandleFunc("/", Index)
	http.HandleFunc("/select/", SelectHandler)
	http.HandleFunc("/create/", CreateHandler)
	http.HandleFunc("/view/", ViewHandler)
	http.HandleFunc("/search/", SearchHandler)
}
