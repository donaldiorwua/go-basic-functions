package gobasicfunctions

import (
	"strconv"
	"strings"
)

func BaseConversionCommand(word []string) []string {
	var base int
	var result []string

	for _, words := range word {
		switch words {
		case "(oct)":
			base = 8
		case "(bin)":
			base = 2
		default:
			base = 16
		}
		switch words {
		case "(hex)", "(oct)", "(bin)":
			index := len(result) - 1
			num, err := strconv.ParseInt(result[index], base, 64)
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(num, 10)
			}
		case "(up)":
			result[len(result)-1] = strings.ToUpper(result[len(result)-1])
		case "(cap)":
			result[len(result)-1] = strings.Title(result[len(result)-1])
		case "(low)":
			result[len(result)-1] = strings.ToLower(result[len(result)-1])
		default:
			result = append(result, words)

		}
	}
	return result
}
