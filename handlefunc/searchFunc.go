package handlefunc

import (
	"net/http"
	"strconv"
	"strings"
)

//SearchHandler Used to generate search page based on the results of the select page
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
		_, studentArray := APICallStudent(w, r, "/students/search", "GET", value)

		searchResult += "<tr> <th> ID </th> <th> Name </th> </tr>"
		for i := range studentArray {

			searchResult += "<tr>"
			searchResult += "<td>" + strconv.Itoa(studentArray[i].ID) + "</td>"
			searchResult += "<td>" + studentArray[i].Name + "</td>"
			searchResult += "</tr>"

		}

		data := Page{Title: "search Page", Content: searchResult, Class: true}
		TemplateInit(w, "searchResults.html", data)

	case "result/teacher":

		data := Page{Title: "search Page", Content: "Test", Class: true}
		TemplateInit(w, "searchResult.html", data)

	case "result/class":

		data := Page{Title: "search Page", Content: "Test", Class: true}
		TemplateInit(w, "searchResult.html", data)

	}

}
