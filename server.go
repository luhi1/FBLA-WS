package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs.ServeHTTP())
	//Create handler for test.html to run the file and shit
	fmt.Println("Server is running on port 8080")

	// Start server on port specified above
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
