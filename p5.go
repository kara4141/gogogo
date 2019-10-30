package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Let's Go!")
}

// func about_handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "About Go")
// }

func main() {
	http.HandleFunc("/", index_handler)
	// http.HandleFunc("/about", index_handler)
	http.ListenAndServe(":8000", nil)
}
