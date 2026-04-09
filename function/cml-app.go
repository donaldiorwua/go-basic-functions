package gobasicfunctions

import (
	"fmt"
	"strings"
)

func PerformOperation() {
	fmt.Println("=========================================")
	fmt.Println("---WELCOME TO COMMANDLINE TEXT EDITOR---")
	fmt.Println("==========================================")

	for {
		var operation, words string

		fmt.Println("Enter a word to transform or type 'exit' to exit the program!")
		fmt.Scanln(&words)

		if words == "exit" {
			fmt.Println("Goodbye! Thank you for using our Editor")
			return
		}

		fmt.Println("Choose an operation: \n lastchar\n capitalize\n deleteindex")
		fmt.Scanln(&operation)

		switch operation {
		case "lastchar":
			fmt.Println("Result: ", lastChar(words))

		case "capitalize":
			fmt.Println("Result: ", capitalizeWord(words))

		case "deleteindex":
			var num int
			fmt.Println("Enter an Index to Delete:")
			fmt.Scanln(&num)
			fmt.Println("Result: ", deleteIndex(words, num))

		default:
			fmt.Println("Invalid Operation")
			continue
		}
	}
}

func lastChar(word string) string {
	if len(word) == 0 {
		fmt.Println("Enter a Word!")
		return ""
	}
	return string(word[len(word)-1])
}

func capitalizeWord(word string) string {
	return strings.ToUpper(word)
}

func deleteIndex(word string, index int) string {
	if index < 0 || index > len(word) {
		fmt.Println("Enter a valid number!")
		return word
	}
	return word[:index] + word[index+1:]
}

func exit() {
	fmt.Println("Exiting...")
}
