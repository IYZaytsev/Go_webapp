package handlefunc

import (
	"net/http"
	"strconv"
	"strings"
)

//SearchHandler Used to generate search page based on the results of the select page
//It also generates the search result page
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	//grabs results from select page drop down
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
	//grabs results from URL to createResult page and delivers results with switch
	urlSubString := r.URL.Path[len("/search/"):]

	var searchResult string

	switch urlSubString {
	case "result/student":
		value = r.FormValue("studentName") + "/" + r.FormValue("enrolledClasses") + "/" + r.FormValue("studentID")

		_, studentSlice := APICallStudent(w, r, "/students/search", "GET", value)

		searchResult += "<tr> <th> ID </th> <th> Name </th> <th> ClassEnrolled </th> <th> Update Entry</th> <th> Delete Entry </th> </tr>"
		for i := range studentSlice {

			searchResult += "<tr>"
			searchResult += "<td>" + strconv.Itoa(studentSlice[i].ID) + "</td>"
			searchResult += "<td>" + studentSlice[i].Name + "</td>"
			searchResult += "<td>"
			for z := range studentSlice[i].ClassList {
				searchResult += studentSlice[i].ClassList[z].Name
			}
			searchResult += "</td>"
			searchResult += "<td>" + "<a href = /update/student/" + strconv.Itoa(studentSlice[i].ID) + ">" + "Update Entry" + "</a>" + "</td>"
			searchResult += "<td>" + "<a href = /delete/student/" + strconv.Itoa(studentSlice[i].ID) + ">" + "Delete Entry" + "</a>" + "</td>"
			searchResult += "</tr>"

		}

		data := Page{Title: "search Page", Content: searchResult, Student: true}
		TemplateInit(w, "searchResults.html", data)

	case "result/teacher":

		value = r.FormValue("teacherName") + "/" + r.FormValue("classes_Assigned") + "/" + r.FormValue("teacherID")

		_, teacherSlice := APICallTeacher(w, r, "/teachers/search", "GET", value)

		searchResult += "<tr> <th> ID </th> <th> Name </th> <th> Update Entry</th> <th> Delete Entry </th> </tr>"
		for i := range teacherSlice {

			searchResult += "<tr>"
			searchResult += "<td>" + strconv.Itoa(teacherSlice[i].ID) + "</td>"
			searchResult += "<td>" + teacherSlice[i].Name + "</td>"
			searchResult += "<td>" + "<a href = /update/teacher/" + strconv.Itoa(teacherSlice[i].ID) + ">" + "Update Entry" + "</a>" + "</td>"
			searchResult += "<td>" + "<a href = /delete/teacher/" + strconv.Itoa(teacherSlice[i].ID) + ">" + "Delete Entry" + "</a>" + "</td>"
			searchResult += "</tr>"

		}

		data := Page{Title: "search Page", Content: searchResult, Teacher: true}
		TemplateInit(w, "searchResults.html", data)

	case "result/class":

		value = r.FormValue("className") + "/" + r.FormValue("classTeacher") + "/" + r.FormValue("classID")

		_, classSlice := APICallClass(w, r, "/classes/search", "GET", value)

		searchResult += "<tr> <th> ID </th> <th> Name </th> <th> TeacherName </th> <th> Update Entry</th> <th> Delete Entry </th></tr>"
		for i := range classSlice {

			searchResult += "<tr>"
			searchResult += "<td>" + strconv.Itoa(classSlice[i].ID) + "</td>"
			searchResult += "<td>" + classSlice[i].Name + "</td>"
			searchResult += "<td>" + classSlice[i].AssignedTeacher.Name + "</td>"
			searchResult += "<td>" + "<a href = /update/class/" + strconv.Itoa(classSlice[i].ID) + ">" + "Update Entry" + "</a>" + "</td>"
			searchResult += "<td>" + "<a href = /delete/class/" + strconv.Itoa(classSlice[i].ID) + ">" + "Delete Entry" + "</a>" + "</td>"
			searchResult += "</tr>"

		}

		data := Page{Title: "search Page", Content: searchResult, Class: true}
		TemplateInit(w, "searchResults.html", data)

	}

}
