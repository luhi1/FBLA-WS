package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("/joe", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Biden")
	})

	fmt.Println("Server is running on port 8080")

	// Start server on port specified above
	http.ListenAndServe(":8080", nil)
}
