package gobasicfunctions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StringManipulator(word []string) []string {
	var base int
	var result []string
	for _, words := range word {
		index := len(result) - 1
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
			num, err := strconv.ParseInt(result[index], base, 64)
			if err == nil {
				result[index] = strconv.FormatInt(num, 10)
			}
		case "(up)":
			result[index] = strings.ToUpper(result[index])
		case "(cap)":
			result[index] = strings.ToUpper(string(result[index][0])) + strings.ToLower(result[index][1:])
		case "(low)":
			result[index] = strings.ToLower(result[index])
		default:
			result = append(result, words)

		}
	}
	return result
}

func maincont() {
	if len(os.Args) < 3 {
		fmt.Println("error")
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]

	input, _ := os.ReadFile(inputfile)

	result := processor(string(input))

	err := os.WriteFile(outputfile, []byte(result), 0644)
	if err == nil {
		fmt.Println("sucess")
	}
}

func processor(text string) string {
	text = textTransformer(text)
	text = punctuation(text)
	//text = article(text)
	text = upNwords(text)

	return text
}

func textTransformer(text string) string {
	words := strings.Fields(text)

	for i := range words {
		if words[i] == "(cap)" {
			words[i-1] = strings.ToUpper(string(words[i-1][0])) + strings.ToLower(words[i-1][1:])
			words[i] = ""
		} else if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(string(words[i-1]))
			words[i] = ""
		} else if words[i] == "(low)" {
			words[i-1] = strings.ToLower(string(words[i-1]))
			words[i] = ""
		} else if words[i] == "(hex)" {
			num, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				fmt.Println(err)
			}
			words[i-1] = strconv.FormatInt(num, 10)
			words[i] = ""
		} else if words[i] == "(bin)" {
			num, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				fmt.Println(err)
			}
			words[i-1] = strconv.FormatInt(num, 10)
			words[i] = ""
		}
	}
	return strings.Join(words, " ")
}

func punctuation(text string) string {

	text = strings.ReplaceAll(text, " '", "'")
	text = strings.ReplaceAll(text, "' ", " '")
	text = strings.ReplaceAll(text, " ,", ", ")
	text = strings.ReplaceAll(text, " !", "!")
	text = strings.ReplaceAll(text, " ?", "!")
	text = strings.ReplaceAll(text, " ...", "...")

	return text
}

func article(text string) string {
	words := strings.Fields(text)
	vow := "aoiuehAOIUE"

	for i := range words {
		if strings.HasPrefix(words[i], vow) {
			words[i-1] = "an"
		} else {
			words[i-1] = "An"
		}

	}
	return strings.Join(words, " ")
}

func upNwords(text string) string {
	wrd := strings.ReplaceAll(text, ", ", ",")
	words := strings.Fields(wrd)
	var result []string

	for _, word := range words {
		if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") {
			trimed := strings.Trim(word, "()")
			splited := strings.Split(strings.TrimSpace(trimed), ",")
			trigerword := strings.TrimSpace(splited[0])
			fmt.Println(trigerword)
			count := 1
			if len(splited) > 1 {
				num, err := strconv.Atoi(splited[1])
				fmt.Println(num)
				if err != nil {
					fmt.Println(err)
				}
				count = num
				fmt.Println(count)
			}
			start := len(result) - count
			if start < 0 {
				start = 0
			}
			for i := start; i < len(result); i++ {
				if trigerword == "up" {
					result[i] = strings.ToUpper(result[i])
				}
			}
		} else {
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}
