package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tplJoe = template.Must(template.ParseFiles("joe.html"))
var tplIndex = template.Must(template.ParseFiles("index.html"))

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/joe", joeHandler)
	http.HandleFunc("/", indexHandler)

	//Start server run, files, and other shit.
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func joeHandler(w http.ResponseWriter, r *http.Request) {
	tplJoe.Execute(w, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tplIndex.Execute(w, nil)
}

/*
	Add error handling and checks to make sure you on the right url (index) before handling.
*/
