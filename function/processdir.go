package gobasicfunctions

import (
	"fmt"
	"os"
	"path/filepath"
)

func ProcessDir() {
	files, err := filepath.Glob("*.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	result, err := os.OpenFile("summary.log", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer result.Close()

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		_, err = result.WriteString(string(content) + "\n")
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
	fmt.Println("Summary written successfully")
}
