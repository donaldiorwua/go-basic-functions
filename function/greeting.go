package gobasicfunctions

import "fmt"

func GreetingText() {
	var input, mgs string
	fmt.Println("Enter Your Name: ")
	fmt.Scanln(&input)
	mgs = "welcome to learn2earn, where learning is turned into earning"
	fmt.Println("Hello " + input + ", " + mgs)
}
