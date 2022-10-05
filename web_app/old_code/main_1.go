package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			log.Panic("Error: ", err)
		}
		fmt.Println(fmt.Sprintf("Number of byter wirtten: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}
