package gobasicfunctions

import (
	"strconv"
	"strings"
)

func HandleCapN(text string) string {
	var result []string
	words := strings.Fields(text)
	for i, word := range words {
		trimed := strings.Trim(word, "()")
		normalized := strings.ToLower(trimed)
		if normalized == "cap" {
			numStr := strings.TrimSuffix((words[i+1]), ")")
			num, err := strconv.Atoi(numStr)
			words[i+1] = ""

			if err != nil {
				continue
			}

			for count := 1; count <= num && i-count >= 0; count++ {
				result[i-count] = strings.ToUpper(string(result[i-count][0])) + strings.ToLower(result[i-count][1:])
			}
			continue
		}
		result = append(result, word)
	}
	return strings.Join(result, " ")
}
