package gobasicfunctions

import "strings"

func CapitalizeLastTwo(words []string, n int) string {
	for i := len(words) - n; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return strings.Join(words, " ")
}
