package gobasicfunctions

import "strings"

func CapitaliseIndexCommand(text string) string {
	words := ManualFields(text)
	for i := range words {
		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words[i] = ""
		}
	}
	return strings.Join(words, " ")
}