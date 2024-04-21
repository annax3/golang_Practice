package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello Prashant, this is your Go server!")
	})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		return
	}
}
