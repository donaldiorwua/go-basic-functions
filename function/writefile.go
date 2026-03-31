package gobasicfunctions

import (
	"os"
)

func WriteFile(words string) {
	filename := "sample.txt"
	data := "Hello There, welcome to Go Lang!"

	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return
	}
}
