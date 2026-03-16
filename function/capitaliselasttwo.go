package gobasicfunctions

import "strings"

func CapitalizeLastTwo(words []string) []string {
	for i := max(0, len(words)-2); i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return words
}
