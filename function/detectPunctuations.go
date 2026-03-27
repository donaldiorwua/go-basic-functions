package gobasicfunctions

import "strings"

func Punc(s string) bool {
	return strings.ContainsAny(s, ",.?;:!")
}
