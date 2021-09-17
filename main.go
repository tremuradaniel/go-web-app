package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	division, err := divideValues(2.0, 0)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("%s", err))
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("This is the division page and x / y = %f", division))

}

func divideValues(x, y float64) (float64, error) {

	if y == 0 {
		err := errors.New("cannot divide by 0!")
		return 0, err
	}

	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// _ = means that if there is an error, pay no attention to it
	_ = http.ListenAndServe(portNumber, nil)
}
