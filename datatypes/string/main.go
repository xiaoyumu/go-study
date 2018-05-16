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

	subStringIndex := strings.Index("abcdef", "cde")

	// fmt.Sprintf(format, args) equal to string.Format(format, args) in C#
	fmt.Println(fmt.Sprintf("The index of string 'cde' in 'abcdef' is %d", subStringIndex))

	const useageMessage = `Usage:
	command [options] [flags]`

	fmt.Println(useageMessage)
	fmt.Println(addThousandSeperator("1233456789"))
}

// Append thousand seperators into the value string
func addThousandSeperator(value string) string {
	length := len(value)
	if length <= 3 {
		return value
	}
	return addThousandSeperator(value[:length-3]) + "," + value[length-3:]
}
