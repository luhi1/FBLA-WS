package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	
	fmt.Println("Server is running on port 8080")

	// Start server on port specified above
	http.ListenAndServe(":8080", nil)
}
