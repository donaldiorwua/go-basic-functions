package gobasicfunctions

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	text, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Unable reading file", err)
		return
	}
	new_text := strings.ToUpper(string(text))

	err = os.WriteFile(outputFile, []byte(new_text), 0644)
	if err != nil {
		fmt.Println("File not found", err)
	}
	fmt.Println("Operation completed successfully! Good job!!!")
}
