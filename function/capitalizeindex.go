package gobasicfunctions

import "strings"

func CapitaliseIndex(word string) string {
	for len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}
