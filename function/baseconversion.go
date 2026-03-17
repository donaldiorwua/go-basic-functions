package gobasicfunctions

import (
	"strconv"
)

func BaseTenConversion(words string, base int) (int64, error) {
	result, err := strconv.ParseInt(words, base, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
