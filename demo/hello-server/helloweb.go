package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

func helloHandler(writer http.ResponseWriter, requuest *http.Request) {
	// WriteHeader() function must be called before writing any content in the writer
	// otherwies you will get a message like yyyy/MM/dd HH:mm:ss http: multiple response.WriteHeader calls
	// it because when calling fmt.Fprintf(writer,"", args), it will call writer.Write() function, and if there is no calls
	// on response.WriteHeader() function before writer.Write() is called, the function response.WriteHeader()
	// will be called with status code 200 automatically.
	writer.WriteHeader(201)
	fmt.Printf("    Request received [Path:%s URI:%s]\n", requuest.URL.Path, requuest.RequestURI)
	fmt.Fprintf(writer, "Hello world, %s", requuest.URL.Path[1:])

}

func now() string {
	return time.Now().Format(time.RFC3339)
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Printf("[%s] >>>Hanlder function [%s] call begin\n", now(), name)
		h(w, r)
		fmt.Printf("[%s] <<<Hanlder function [%s] call finished\n", now(), name)
	}
}

func main() {
	port := "8080"
	fmt.Printf("Starting server (Port: %s) ... Press Ctrl + C to stop.\n", port)
	http.HandleFunc("/hello", log(helloHandler))
	http.ListenAndServe(":"+port, nil)
}
