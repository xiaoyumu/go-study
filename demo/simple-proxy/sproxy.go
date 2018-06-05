package main

import (
	"fmt"
	"net/http"
	"os"
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

func log(handler http.HandlerFunc) http.HandlerFunc {

	name := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("[%s] >>> Begin calling hanlder function [%s]\n", now(), name)
		handler(w, r)
		fmt.Printf("[%s] <<< Finished called hanlder function [%s]\n", now(), name)
	}
}

const DefaultPort string = "8080"

type ProxyParameter struct {
	hostPort  string
	aliasURL  string
	targetURL string
}

func getParameter(args []string) ProxyParameter {
	for _, arg := range args {
		fmt.Println("P:" + arg)
	}
	return ProxyParameter{
		hostPort: DefaultPort,
	}
}

func main() {
	parameter := getParameter(os.Args)

	fmt.Printf("Starting server (Port: %s) ... Press Ctrl + C to stop.\n", parameter.hostPort)
	http.HandleFunc("/hello", log(helloHandler))
	http.ListenAndServe(":"+parameter.hostPort, nil)
}
