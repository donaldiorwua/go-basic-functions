package gobasicfunctions

import "strings"

func FixSentenceArticles(text string) string {
	words := strings.Fields(text)
	//for i := 0; i < len(words); i++ {
	for i := range words {
		if strings.ToLower(words[i]) == "a" || strings.ToLower(words[i]) == "o" {
			words[i] = "an"
		}
	}
	return strings.Join(words, " ")
}
