package main

import (
	"fmt"
	gobasicfunctions "go-basic-functions/function"
)

func main() {

	// words := []string{"this", "is", "a", "simple", "test", "' awesome '"}

	// fmt.Println(gobasicfunctions.CapitalizeLastTwo(words, 2))
	// fmt.Println(gobasicfunctions.BaseTenConversion("FF", 16))
	// fmt.Println(gobasicfunctions.BaseTenConversion("1010", 2))
	// fmt.Println(gobasicfunctions.BaseTenConversion("37", 8))
	// fmt.Println(gobasicfunctions.BaseTenConversion("99", 1))
	// fmt.Println(gobasicfunctions.CapitaliseIndex("hello"))
	// gobasicfunctions.GreetingText()
	// fmt.Println(gobasicfunctions.Punc("m"))
	// fmt.Println(gobasicfunctions.FixSentenceArticles("There is a apple"))
	//fmt.Println(gobasicfunctions.FixSentenceArticles("a egg"))
	//fmt.Println(gobasicfunctions.RemoveComm([]string{"go", "to", "school", "(up)"}, 3))
	// fmt.Println(gobasicfunctions.ReverseString("hello"))
	//gobasicfunctions.PerformOperation()
	//fmt.Println(gobasicfunctions.CountWords("go go is so so fun"))
	//Data := ""
	//file := ""
	//gobasicfunctions.WriteFile(Data)
	//gobasicfunctions.ReadFile(file)
	//gobasicfunctions.ProcessDir()
	// fmt.Println(gobasicfunctions.StringManipulator([]string{"hello", "111", "(bin)", "world"}))
	// fmt.Println(gobasicfunctions.StringManipulator([]string{"hello", "1e", "(hex)", "world"}))
	// fmt.Println(gobasicfunctions.StringManipulator([]string{"hello", "145", "(oct)", "world"}))
	fmt.Println(gobasicfunctions.StringManipulator([]string{"hello", "welcome", "(up)", "world"}))
	fmt.Println(gobasicfunctions.StringManipulator([]string{"hello", "welcome", "(cap)", "world"}))
	fmt.Println(gobasicfunctions.StringManipulator([]string{"hello", "WELCOME", "(low)", "world"}))
	fmt.Println(gobasicfunctions.BaseConversionCommand([]string{"hello", "111", "(bin)", "world"}))
	fmt.Println(gobasicfunctions.BaseConversionCommand([]string{"hello", "1e", "(hex)", "world"}))
	fmt.Println(gobasicfunctions.BaseConversionCommand([]string{"hello", "145", "(oct)", "world"}))
}
