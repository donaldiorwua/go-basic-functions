package gobasicfunctions

import (
	"fmt"
	"strings"
)

func CountWords(str string) string {
	words := strings.Split(str, " ")
	count := make(map[string]int)

	for _, word := range words {
		count[word]++
	}
	return fmt.Sprint(count)
}
