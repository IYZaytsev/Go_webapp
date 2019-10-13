package handlefunc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"

	"../databasetypes"
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

var myClient = http.Client{Timeout: 10 * time.Second}
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

//APICallStudent is a multi use api call that returns a student struct
func APICallStudent(w http.ResponseWriter, r *http.Request, path string, method string, parameter string) (databasetypes.Student, databasetypes.Students) {
	urlString := "http://localhost:8080" + path
	studentStruct := databasetypes.Student{}
	studentArray := databasetypes.Students{}
	req, err := http.NewRequest(method, urlString, bytes.NewBufferString(parameter))
	if err != nil {
		log.Fatal(err)
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	res, getErr := myClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	if method == "GET" {
		jsonErr := json.Unmarshal(body, &studentArray)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
	} else {
		jsonErr := json.Unmarshal(body, &studentStruct)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
	}

	return studentStruct, studentArray
}

//APICallTeacher is a multi use api call that returns a teacher struct
func APICallTeacher(w http.ResponseWriter, r *http.Request, path string, method string, parameter string) databasetypes.Teacher {
	urlString := "http://localhost:8080" + path

	req, err := http.NewRequest(method, urlString, bytes.NewBufferString(parameter))
	if err != nil {
		log.Fatal(err)
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	res, getErr := myClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	teacherStruct := databasetypes.Teacher{}
	jsonErr := json.Unmarshal(body, &teacherStruct)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return teacherStruct
}

//APICallClass is a multi use api call that returns a teacher struct
func APICallClass(w http.ResponseWriter, r *http.Request, path string, method string, parameter string) databasetypes.Class {
	urlString := "http://localhost:8080" + path

	req, err := http.NewRequest(method, urlString, bytes.NewBufferString(parameter))
	if err != nil {
		log.Fatal(err)
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	res, getErr := myClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	classStruct := databasetypes.Class{}
	jsonErr := json.Unmarshal(body, &classStruct)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return classStruct
}
