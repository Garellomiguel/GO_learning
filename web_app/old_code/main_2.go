package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// For a func to response a request have to handle two parameters
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Web Home Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum, _ := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is a sum of 2 and 2: %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(2, 0)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error: %d", err))
		return
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is a division of %d and %d: %d", f, f, f))
}

func addValues(x int, y int) (int, error) {
	return x + y, nil
}

func divideValues(x float32, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("Cannot divide by 0")
		return 0, err
	} else {
		result := x / y
		return result, nil
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %v \n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
