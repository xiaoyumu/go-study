package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	calcFib()
}

// The trace function will be used with function defer
func traceFunctionExecution(name string) func() {
	start := time.Now()
	log.Printf("Start %s", name)
	return func() { log.Printf("%s finished (Toke %s)", name, time.Since(start)) }
}

func calcFib() {
	// Trace execution start and defer the finish call.
	defer traceFunctionExecution("calcFib")()
	fmt.Printf("Calculating ...  ")
	// Show spinner in a go routine
	go spinner(100 * time.Millisecond)
	const n = 45
	result := fib(n)
	fmt.Printf("\nFibonacci(%d) = %d\n", n, result)
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			// \b will backspace 1 char before print out the char spinner char
			fmt.Fprintf(os.Stdout, "\b%c", r)
			time.Sleep(delay)
		}
	}
}
