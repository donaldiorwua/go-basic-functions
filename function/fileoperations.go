package gobasicfunctions

import (
	"fmt"
	"io"
	"os"
)


func filemain() {
	//file opening
	file, err := os.Open("sample.txt")
	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()

	// reading file content
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(content))

	//file creation
	newfile, err := os.Create("outputfile.txt")
	if err != nil{
		fmt.Println(err)
	}
	defer newfile.Close()

	//writing to a file
	data := "I am a little bit opssessed"

	_, err = newfile.WriteString(data)
	if err != nil{
		fmt.Println(err ) 
	}

	//opening file with flags using os.OpenFile()
	file1, err := os.OpenFile("outputfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()

	file1.WriteString("I have added new text")
}