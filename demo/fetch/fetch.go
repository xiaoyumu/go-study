package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// The first element in os.Args is the full path of the program itself.
	// So we start to pick args from os.Args[1]
	for _, url := range os.Args[1:] {
		fmt.Println("Requesting : " + url)

		// According to the implemention of http.Get(url), when err != nil
		// the response is nil.
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatch url %s failed due to %v.\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("Response Code: %d\n", response.StatusCode)
		bytes, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Faild to read response content due to %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", bytes)
	}
}
