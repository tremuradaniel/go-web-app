package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8082"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the main page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// _ = means that if there is an error, pay no attention to it
	_ = http.ListenAndServe(portNumber, nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}
