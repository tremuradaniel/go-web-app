package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8082"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// _ = means that if there is an error, pay no attention to it
	_ = http.ListenAndServe(portNumber, nil)
}
