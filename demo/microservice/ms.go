package main

import (
	"fmt"
	"os"

	"github.com/xiaoyumu/go-study/commandline"
)

// DefaultHost for http server to listen
const DefaultHost string = "localhost"

// DefaultPort for http server to listen
const DefaultPort string = "8000"

func main() {
	parameterParseSetting := &commandline.ParameterParseSetting{
		RequiredParameters: []string{"host", "port"},
	}
	pool, err := commandline.Create(parameterParseSetting)
	if err != nil {
		fmt.Println(err)
		displayUsage()
		os.Exit(-1)
	}

	pool.DumpParameters()
}

func displayUsage() {
	fmt.Println()
	fmt.Println("Usage: ")
	fmt.Println("    executable -host:<Host|IP> -port:<8080>")
}
