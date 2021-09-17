package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8082"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the main page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, "This is the about page and 2 + 2 = %d", sum)
}

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// _ = means that if there is an error, pay no attention to it
	_ = http.ListenAndServe(portNumber, nil)
}
