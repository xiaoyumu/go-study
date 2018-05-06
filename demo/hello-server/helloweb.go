package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, requuest *http.Request) {
	fmt.Printf("Request received.\n")
	fmt.Fprintf(writer, "Hello world, %s", requuest.URL.Path[1:])
}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
