package main

import (
	gobasicfunctions "go-basic-functions/function"
	"strings"
)

func capitalise(words []string) []string {
	for i := len(words) - 2; i < len(words); i++ {
		if len(words) >= 0 {
			words[i] = strings.ToUpper(words[i])
		}
	}
	return words
}

func main() {

	/*words := []string{"this", "is", "a", "simple", "test", "' awesome '"}

	result := capitalise(words)

	fmt.Println(capitalise(result))*/
	/*fmt.Println(baseTenConversion("FF", 16))
	fmt.Println(baseTenConversion("1010", 2))
	fmt.Println(baseTenConversion("37", 8))
	fmt.Println(baseTenConversion("99", 1))*/
	//fmt.Println(CapitaliseIndex("hello"))
	gobasicfunctions.GreetingText()

}

func CapitaliseIndex(word string) string {
	/*for len(word) == 0 {
		return word
	}*/
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}
