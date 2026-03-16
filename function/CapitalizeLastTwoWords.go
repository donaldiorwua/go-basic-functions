package gobasicfunctions

import "strings"

func CapitalizeLastTwoWords(words []string) []string {
	if len(words) < 2 {
		return words
	}
	words[len(words)-1] = strings.ToUpper(words[len(words)-1])
	words[len(words)-2] = strings.ToUpper(words[len(words)-2])

	return words
}
