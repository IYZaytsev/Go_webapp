package handlefunc

import (
	"fmt"
	"net/http"
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
	switch urlSubString {
	case "result/student":
		studentid := test(w, r)
		contentstring := fmt.Sprintf("student id %d", studentid)
		data := Page{Title: "search Page", Content: contentstring, Class: true, Type: "class"}
		TemplateInit(w, "search.html", data)

	}

}
