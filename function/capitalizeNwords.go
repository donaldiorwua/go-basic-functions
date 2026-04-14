package gobasicfunctions

import (
	"strconv"
	"strings"
)

func HandleCapN(text string) string {
	var result []string
	words := strings.Fields(text)
	for _, word := range words {
		if word == "(up,2)" {
			trimed := strings.Trim(word, "()")
			normalized := strings.ToLower(trimed)

			splitted := strings.Split(normalized, ",")
			removespace := strings.TrimSpace(splitted[0])

			count := 1
			if len(splitted) > 1 {
				num, _ := strconv.Atoi(strings.TrimSpace(splitted[1]))
				if num > count {
					count = num
				}
			}
			start := len(result) - count
			if count > start {
				start = 0
			}
			for k := start; k < len(result); k++ {
				if removespace == "up" {
					result[k] = strings.ToUpper(result[k])
				}
			}
			continue
		}
		result = append(result, word)
	}
	return strings.Join(result, " ")
}
