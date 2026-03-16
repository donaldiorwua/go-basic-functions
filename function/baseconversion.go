package gobasicfunctions

import (
	"strconv"
)

func BaseConversion(word string, base int) (int64, error) {
	result, _ := strconv.ParseInt(word, base, 64)
	return result, nil
}
