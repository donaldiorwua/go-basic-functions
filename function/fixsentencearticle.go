package gobasicfunctions

import "strings"

func FixSentenceArticles(text string) string {
	words := strings.Fields(text)
	for i := 1; i < len(words); i++ {
		if strings.ToLower(words[i]) == "a" {
			words[i] = "an"
		}
	}
	return strings.Join(words, " ")
}
