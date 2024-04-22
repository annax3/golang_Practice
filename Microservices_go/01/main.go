package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, err2 := fmt.Fprintf(writer, "Hello Prashant, this is your Go server!")
		if err2 != nil {
			writer.WriteHeader(http.StatusBadRequest)
		}
		data, err := io.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("OOPS"))
			return
		}
		fmt.Printf("Data  %s ", data)
	})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		return
	}
}
