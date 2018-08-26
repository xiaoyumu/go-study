package main

import (
	"log"
	"time"
)

func main() {

	nautuals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 1000; x++ {
			nautuals <- x
			time.Sleep(500)
		}
		close(nautuals)
	}()

	go func() {
		for {
			x, ok := <-nautuals
			if !ok {
				break
			}
			squares <- x * x
			time.Sleep(500)
		}
		close(squares)
	}()

	for v := range squares {
		log.Printf("Square: %v", v)
		time.Sleep(200)
	}
}
