package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"strconv"
)

type UserInformation struct {
	name          string
	grade         uint8
	studentNumber uint32
	password      string
}

//As the project develops, determine whether this struct is required.
//If not, merge into TemplateRequests struct

type TemplateRequests struct {
	data     UserInformation
	filename string
}

var tplErr = template.Must(template.ParseFiles("404.html"))

// Start server run, files, and other shit.
func main() {
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/home", homeHandler)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func tplExec(w http.ResponseWriter, request TemplateRequests) {
	temp := template.Must(template.ParseFiles(request.filename))

	err := temp.Execute(w, request.data)
	if err != nil {
		return
	}
}

func errorTplExec(w http.ResponseWriter) {
	err := tplErr.Execute(w, nil)
	if err != nil {
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fullRequest := TemplateRequests{UserInformation{}, "login.html"}

	if r.URL.String() == "/" {
		tplExec(w, fullRequest)
	} else {
		errorTplExec(w)
	}
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fullRequest := TemplateRequests{
		UserInformation{},
		"signup.html",
	}

	if r.URL.String() == "/signup" {
		tplExec(w, fullRequest)
	} else {
		errorTplExec(w)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fullRequest := TemplateRequests{UserInformation{}, "login.html"}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			return
		}
		idNumber, err := strconv.Atoi(r.FormValue("studentNumber"))

		if err != nil {
			return
		}

		fullRequest = TemplateRequests{
			data: UserInformation{
				name:          "Michael",
				grade:         10,
				studentNumber: uint32(idNumber),
				password:      hashPswd(r.FormValue("password")),
			},
			filename: "home.html",
		}
		//Queries go here

		//TODO: error is with rendering the template, the parsing of the form itself works
	}

	if r.URL.String() == "/home" {
		tplExec(w, fullRequest)
	} else {
		errorTplExec(w)
	}

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
