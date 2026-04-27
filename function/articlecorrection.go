package gobasicfunctions

import "strings"

func ArticleCorrection(text string) string {
	words := strings.Fields(text)
	vow := "aoiueAOIUE"
	punc := "'()\":;,.?!"

	for i := range words {
		if words[i] == "a" || words[i] == "A" || words[i] == "an" || words[i] == "An" {
			removepunc := strings.Trim(strings.Join(words[i+1:], " "), punc)
			if strings.ContainsAny(vow, string(removepunc[0])) {
				if words[i] == "a" {
					words[i] = "an"
				} else if words[i] == "A" {
					words[i] = "An"
				}
			} else {
				if words[i] == "an" {
					words[i] = "a"
				} else if words[i] == "An" {
					words[i] = "A"
				}
			}
		}

	}
	return strings.Join(words, " ")
}
