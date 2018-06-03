package main

import "fmt"

func forLoopAsWhile() {
	continueLooping := true

	const MaxCount int = 10
	counter := 0
	fmt.Printf("For looping %d to %d\n", counter, MaxCount)
	for continueLooping {
		if counter < MaxCount {
			fmt.Printf("    At %d\n", counter)
			counter++
		} else {
			continueLooping = false
		}
	}
	fmt.Println("For Loop finished")
}

func forLoopAsIteration() {

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("For looping on slice\n")
	for index, value := range list {
		fmt.Printf("    Value %d at %d\n", value, index)
	}
	fmt.Println("For Loop finished")
}

func forLoopAsIterationForIndexOnly() {

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("For looping on slice\n")
	for index := range list {
		fmt.Printf("    Value %d at %d\n", list[index], index)
	}
	fmt.Println("For Loop finished")
}
