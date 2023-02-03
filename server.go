package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

var tplJoe = template.Must(template.ParseFiles("joe.html"))
var tplIndex = template.Must(template.ParseFiles("index.html"))
var errIndex = template.Must(template.ParseFiles("404.html"))

type JoeyB struct {
	Password string
	Hash     string
}

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

		err := tplJoe.Execute(w, JoeyB{"WhiteHouse123", hashPswd("WhiteHouse123")})
		if err != nil {
			return
		}

	default:
		errIndex.Execute(w, nil)

	}
}

func hashPswd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

/*
	TODO: Seperate the Handlers for each index.
	TODO: Abstract many of the proccesses in functions + variables (less inline calls and declarations)
*/
