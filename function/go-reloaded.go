package gobasicfunctions

import (
	"strconv"
	"strings"
)

func StringManipulator(word []string) []string {
	var base int
	var result []string
	//index := len(result) - 1
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
			num, err := strconv.ParseInt(result[len(result)-1], base, 64)
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

func textTransformer(words []string) []string {
	var result []string
	for _, word := range words {
		switch word {
		case "(up)":

			result[len(result)-1] = strings.ToUpper(result[len(result)-1])

		case "(cap)":
			result[len(result)-1] = strings.Title(result[len(result)-1])

		case "(low)":
			result[len(result)-1] = strings.ToLower(result[len(result)-1])
		default:
			result = append(result, word)
		}
	}
	return result
}

func ToUpperN(word []string) []string {
	for i := len(word) - 1; i > 0; i++ {

	}
	return word
}
