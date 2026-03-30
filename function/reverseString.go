package gobasicfunctions

import (
	"slices"
	"strings"
)

func ReverseString(str string) string {
	word := strings.Split(str, "")
	slices.Reverse(word)
	return strings.Join(word, "")
}
