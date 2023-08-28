package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			fmt.Fprintf(w, "index!!")
		} else if r.URL.Path == "/bonjour" {
			fmt.Fprintf(w, "Bonjour!")
		} else if r.URL.Path == "/hello" {
			fmt.Fprintf(w, "Hello World!")
		} else {
			fmt.Fprintf(w, "404 Not Found")
		}
	})
	http.ListenAndServe(":8888", nil)
}
