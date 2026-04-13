package gobasicfunctions

import "strings"

func Handlecase(text string) string {
	var result []string
	words := strings.Fields(text)

	for _, word := range words {
		index := len(result) - 1
		normalized := strings.ToLower(word)
		switch normalized {
		case "(up)":
			result[index] = strings.ToUpper(result[index])
		case "(low)":
			result[index] = strings.ToLower(result[index])
		case "(cap)":
			result[index] = strings.ToUpper(string(result[index][0])) + strings.ToLower(result[index][1:])
		default:
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}
