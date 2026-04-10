package gobasicfunctions

import (
	"strings"
)

// func StringManipulator(word []string) []string {
// 	var base int
// 	var result []string
// 	for _, words := range word {
// 		index := len(result) - 1
// 		switch words {
// 		case "(oct)":
// 			base = 8
// 		case "(bin)":
// 			base = 2
// 		default:
// 			base = 16
// 		}
// 		switch words {
// 		case "(hex)", "(oct)", "(bin)":
// 			num, err := strconv.ParseInt(result[index], base, 64)
// 			if err == nil {
// 				result[index] = strconv.FormatInt(num, 10)
// 			}
// 		case "(up)":
// 			result[index] = strings.ToUpper(result[index])
// 		case "(cap)":
// 			result[index] = strings.Title(result[index])
// 		case "(low)":
// 			result[index] = strings.ToLower(result[index])
// 		default:
// 			result = append(result, words)

// 		}
// 	}
// 	return result
// }

func textTransformer(words []string) []string {
	var result []string
	index := len(result) - 1
	for _, word := range words {
		switch word {
		case "(up)":

			result[index] = strings.ToUpper(result[index])

		case "(cap)":
			result[index] = strings.Title(result[index])

		case "(low)":
			result[index] = strings.ToLower(result[index])
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
