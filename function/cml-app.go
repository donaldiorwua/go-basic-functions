package gobasicfunctions

import (
	"fmt"
	"strings"
)

func PerformOperation() {
	var words, operation string
	fmt.Println("Enter a Word: ")
	fmt.Scanln(&words)

	fmt.Println("Choose an operation: \n lastChar\n capilizeWord\n deleteIndex")
	fmt.Scanln(&operation)

	switch operation {
	case "lastChar":
		fmt.Println("Result: ", lastChar(words))

	case "capilizeWord":
		fmt.Println("Result: ", capitalizeWord(words))

	case "deleteIndex":
		var num int
		fmt.Println("Enter an Index to Delete:")
		fmt.Scanln(&num)
		defer deleteIndex(words, num)
		fmt.Println("Result: ", deleteIndex(words, num))
	default:
		fmt.Println("Invalid Operation")
	}
}
func lastChar(word string) string {
	if len(word) == 0 {
		fmt.Println("Enter a Word!")
	}
	return string(word[len(word)-1])
}

func capitalizeWord(word string) string {
	return strings.ToUpper(word)
}

func deleteIndex(word string, index int) string {
	if index < 0 || index > len(word) {
		fmt.Println("Enter a valid number!")
	}
	return word[:index] + word[index+1:]
}
