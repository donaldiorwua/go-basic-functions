package gobasicfunctions

import "fmt"

func GreetingText() {
	var input, mgs string
	fmt.Println("Enter Your Name: ")
	fmt.Scanln(&input)
	mgs = "welcome to Go Lang!"
	fmt.Println("Hello " + input + ", " + mgs)
}
