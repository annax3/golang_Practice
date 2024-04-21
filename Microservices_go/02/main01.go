package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		_, err := io.WriteString(w, "Hello From a Handle Function #1!\n")
		if err != nil {
			return
		}
	}

	h2 := func(w http.ResponseWriter, _ *http.Request) {
		_, err := io.WriteString(w, "Hello From a Handle Function #2!\n")
		if err != nil {
			return
		}

	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
