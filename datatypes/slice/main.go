package main

import (
	"fmt"
	"os"
)

func main() {
	intList := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(intList)
	// The following statement equal to
	//   slice := intList[:]
	//   reverse(slice)
	reverse(intList[:])

	// [7 6 5 4 3 2 1]
	fmt.Println(intList)

	intList = [...]int{1, 2, 3, 4, 5, 6, 7}

	annulusMoveToLeft(intList[:], 8)

	fmt.Println(intList)

	// Checking on slice equal
	list1 := []interface{}{1, 2, 3, 4, 5, 6, 7}
	list2 := []interface{}{1, 2, 3, 4, 5, 6, 7}

	if equals(list1, list2) {
		fmt.Println("Slice list1 equals list2")
	}

	list1 = []interface{}{"str1", "str2", "str3"}
	list2 = []interface{}{"str1", "str2", "str4"}

	if equals(list1, list2) {
		fmt.Println("Slice list1 equals list2")
	} else {
		fmt.Println("Slice list1 NOT equals list2")
	}

	element := "str1"
	if contains(list1, element) {
		fmt.Fprintf(os.Stdout, "The element [%v] exists in %v\n", element, list1)
	}

	intElement := 1
	valueList := []interface{}{1, 2, 3, 4, 5}
	if contains(valueList, intElement) {
		fmt.Fprintf(os.Stdout, "The element [%v] exists in %v\n", intElement, valueList)
	}
}

// Reverse a slice
func reverse(values []int) {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
}

// Move the slice in given steps
func annulusMoveToLeft(values []int, steps int) {

	// steps must be less than or equal to the length of the given slice
	if steps > len(values) {
		return
	}

	// If input values are [1 2 3 4 5 6 7]
	// And steps is 2
	// move the [1 2 3 4 5 6 7] << 2
	// it should produce [3 4 5 6 7 1 2]

	reverse(values[:steps])
	// [2 1 3 4 5 6 7]
	//  ^ ^ -- get reversed

	reverse(values[steps:])
	// [2 1 7 6 5 4 3]
	//      ^ ^ ^ ^ ^ -- get reversed

	reverse(values)
	// [3 4 5 6 7 1 2]
	//  ^ ^ ^ ^ ^ ^ ^ -- get reversed
}

// Kinda generic equals implementation
func equals(sliceA []interface{}, sliceB []interface{}) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}

	for i := 0; i < len(sliceA); i++ {
		if sliceA[i] != sliceB[i] {
			return false
		}
	}

	return true
}

// Check if an 'element' exists in the given slice
func contains(slice []interface{}, element interface{}) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
