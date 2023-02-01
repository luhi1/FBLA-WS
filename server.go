package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tplJoe = template.Must(template.ParseFiles("joe.html"))
var tplIndex = template.Must(template.ParseFiles("index.html"))
var errIndex = template.Must(template.ParseFiles("404.html"))

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/joe", Handler)
	http.HandleFunc("/", Handler)

	//Start server run, files, and other shit.
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.String() {

	case "/":
		tplIndex.Execute(w, nil)

	case "/joe":
		tplJoe.Execute(w, nil)

	default:
		errIndex.Execute(w, nil)

	}
}
