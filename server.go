package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"strings"
)

type UserInformation struct {
	Username string
	Hash     string
}

//As the project develops, determine whether this struct is required.
//If not, merge into TemplateRequests struct

type TemplateRequests struct {
	data  UserInformation
	index string
}

var tplErr = template.Must(template.ParseFiles("404.html"))

// Start server run, files, and other shit.
func main() {
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/joe", joeHandler)
	http.HandleFunc("/", indexHandler)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func tplExec(w http.ResponseWriter, r *http.Request, request TemplateRequests) {
	var err error = nil
	temp := template.Must(template.ParseFiles(strings.TrimPrefix(request.index, "/") + ".html"))

	if r.URL.String() == request.index {
		err = temp.Execute(w, request.data)
	} else {
		err = temp.Execute(w, nil)
	}

	if err != nil {
		return
	}
}

func joeHandler(w http.ResponseWriter, r *http.Request) {
	fullRequest := TemplateRequests{
		UserInformation{"abcde", hashPswd("abcde")},
		"/joe",
	}
	tplExec(w, r, fullRequest)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fullRequest := TemplateRequests{
		UserInformation{},
		"/index",
	}
	tplExec(w, r, fullRequest)
}

func hashPswd(pwd string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

/*
	TODO: It's SQL Time :)
*/
