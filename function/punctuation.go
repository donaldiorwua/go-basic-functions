package gobasicfunctions

import "strings"

func IsPun(word string) bool {
	return len(word) == 1 && strings.Contains(".,!?;:'", word)
}
