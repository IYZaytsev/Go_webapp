package main

import (
	"log"
	"net/http"

	"./handlefunc"
)

func main() {

	handlefunc.SetHandlers()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("listening")
	log.Fatalln(http.ListenAndServe(":9080", nil))
}
