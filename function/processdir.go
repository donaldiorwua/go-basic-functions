package gobasicfunctions

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func ProcessDir() {
	files, err := filepath.Glob("*.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		files = append(files, file)
		fmt.Println("Content of", file, ":", string(content))
	}
}
