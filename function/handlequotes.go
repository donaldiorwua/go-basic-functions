package gobasicfunctions

import "strings"

func HandleQuotes(words []string) []string {
	text := strings.Join(words, " ")
	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")

	return strings.Fields(text)
}
