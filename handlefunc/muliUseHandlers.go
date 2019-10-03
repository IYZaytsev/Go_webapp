package handlefunc

import (
	"net/http"
	"text/template"
)

//Page is used with Init function to create HTML pages
type Page struct {
	Title   string
	Header  string
	Action  string
	Teacher bool
	Student bool
	Class   bool
	Type    string
	Content string
	ID      int
}

var tmpls = template.Must(template.ParseFiles("tmpl/index.html", "tmpl/select.html", "tmpl/create.html", "tmpl/view.html",
	"tmpl/search.html", "tmpl/searchResults.html"))

//TemplateInit used every where to initialize HTML templates
func TemplateInit(w http.ResponseWriter, templateFile string, templateData Page) {
	if err := tmpls.ExecuteTemplate(w, templateFile, templateData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//SelectHandler used to generate select page based on the previous page
func SelectHandler(w http.ResponseWriter, r *http.Request) {
	urlSubString := r.URL.Path[len("/select/"):]

	switch urlSubString {
	case "create":
		data := Page{Title: "School Database", Header: "What type would you like to create?", Action: "/" + urlSubString + "/"}
		TemplateInit(w, "select.html", data)

	case "search":
		data := Page{Title: "School Database", Header: "What type would you like to search?", Action: "/" + urlSubString + "/"}
		TemplateInit(w, "select.html", data)

	}

}

//Index is the main page of the webbapp used for naviagation
func Index(w http.ResponseWriter, r *http.Request) {
	data := Page{Title: "School Database", Header: "Welcome, please select an option"}
	TemplateInit(w, "index.html", data)

}
