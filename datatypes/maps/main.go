package main

import "fmt"

func main() {

	// Define and initialize a map (Key as string, Value as int)
	mapStringToInt := map[string]int{
		"Key1": 10,
		"Key2": 20,
		"Key3": 30,
		"Key4": 40,
	}

	for k, v := range mapStringToInt {
		fmt.Printf("Key=%s, Value=%v\n", k, v)
	}

	// Define a map variable
	var map2 map[string]int

	// Initialize it
	map2 = make(map[string]int)

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key-%v", i)
		map2[key] = i * 10
	}

	// you can noticed that keys in the map are not ordered.
	// every time it prints in different order.
	for k, v := range map2 {
		fmt.Printf("Key=%s, Value=%v\n", k, v)
	}

	var x, y int
	x, y = 2, 3
	fmt.Printf("%v * %v = %v\n", x, y, calculate(x, "*", y))

	fmt.Printf("%v mode %v = %v\n", y, x, calculate(y, "%", x))
}

// This map holds calculation functions mapping to corresponding operator
var calculations map[string]func(int, int) int

// Initialize the calculation functions mapping
func init() {
	calculations = make(map[string]func(int, int) int)
	calculations["+"] = add
	calculations["-"] = substract
	calculations["*"] = multiply
	calculations["/"] = divide

	// Anonymous function for mode operation
	calculations["%"] = func(x, y int) int {
		return x % y
	}

}

func calculate(x int, operator string, y int) int {
	var cal func(int, int) int
	if f, exists := calculations[operator]; !exists {
		panic(fmt.Sprintf("The operator [%s] was not supported.", operator))
	} else {
		cal = f
	}

	return cal(x, y)
}

func add(x, y int) int {
	return x + y
}

func substract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
