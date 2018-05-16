package main

import "fmt"
import "strings"

func main() {
	fmt.Printf("%s\n", "a string value")

	var stringValue string
	stringValue = "/home/xiaoyu/dev/test.go"

	parts := strings.Split(stringValue, "/")

	// slice the parts from index 1 to skip the first empty string
	for _, part := range parts[1:] {
		fmt.Println(part)
	}

	rejoinString := strings.Join(parts, "\\")

	fmt.Printf("Rejoined string: %s\n", rejoinString)

	subStringValue := "xiaoyu"

	if strings.Contains(stringValue, subStringValue) {
		fmt.Printf("\"%s\" contains \"%s\"\n", stringValue, subStringValue)
	}

	subStringValue2 := "xiaoyumu"

	if !strings.Contains(stringValue, subStringValue2) {
		fmt.Printf("\"%s\" does not contain \"%s\"\n", stringValue, subStringValue2)
	}
}
