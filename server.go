package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"strconv"
)

// UserInformation Template variable names must be capitalized
type UserInformation struct {
	Name          string
	Grade         uint8
	StudentNumber uint32
	Password      string
	Invalid       bool
}

var userInfo = UserInformation{}
var tplErr = template.Must(template.ParseFiles("404.gohtml"))

// Start server run, files, and other shit.
func main() {
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/logout", logoutHandler)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func tplExec(w http.ResponseWriter, filename string) {
	temp := template.Must(template.ParseFiles(filename))

	err := temp.Execute(w, userInfo)
	if err != nil {
		return
	}
}

// Could use this to display error info later on, not limited to just 404s.
func errorTplExec(w http.ResponseWriter) {
	err := tplErr.Execute(w, nil)
	if err != nil {
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/" {
		if userInfo != (UserInformation{}) && !userInfo.Invalid {
			tplExec(w, "home.gohtml")
		} else {
			tplExec(w, "login.gohtml")
			userInfo.Invalid = false
		}
	} else {
		errorTplExec(w)
	}
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if userInfo != (UserInformation{}) && !userInfo.Invalid {
		http.Redirect(w, r, "./home", 303)
		tplExec(w, "home.gohtml")
	} else {
		tplExec(w, "signup.gohtml")
		//dont trust html for valid for results, check it here too
		//figure out how to do signup once database is connected
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	userInfo = UserInformation{}
	http.Redirect(w, r, "/", 303)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			return
		}

		idNumber, err := strconv.Atoi(r.FormValue("studentNumber"))
		if err != nil {
			return
		}

		//Make sure that the input fits the requirements (the student number is a uint32 < 9999999)
		//Check with query in the future
		if idNumber == 12333 {
			userInfo = UserInformation{
				Name:          "Michael",
				Grade:         10,
				StudentNumber: uint32(idNumber),
				Password:      hashPswd(r.FormValue("password")),
				Invalid:       false,
			}
			tplExec(w, "home.gohtml")
		} else {
			userInfo.Invalid = true
			http.Redirect(w, r, "./", 303)
		}

	} else {
		http.Redirect(w, r, "./", 303)
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
