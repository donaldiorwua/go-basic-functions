package main

import (
	"fmt"
	gobasicfunctions "go-basic-functions/function"
	"strings"
)

/*
func caseConversion(words []string) []string {
	if len(words) < 2 {
		return words
	}
	words[len(words)-1] = strings.ToUpper(words[len(words)-1])
	words[len(words)-2] = strings.ToUpper(words[len(words)-2])

	return words
}*/

func capitalise(words []string) []string {
	for i := len(words) - 2; i < len(words); i++ {
		if len(words) >= 0 {
			words[i] = strings.ToUpper(words[i])
		}
	}
	return words
}

func main() {

	words := []string{"this", "is", "a", "simple", "test", "' awesome '"}

	result := gobasicfunctions.HandleQuotes(words)

	fmt.Println(gobasicfunctions.HandleQuotes(result))

}
