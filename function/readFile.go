package gobasicfunctions

import (
	"fmt"
	"os"
)

func ReadFile(filename string) {
	file := "sample.txt"

	content, err := os.ReadFile(file)
	if err != nil {
		return
	}
	fmt.Println(string(content))
}
