package gobasicfunctions

import (
	"os"
)

func WriteFile(words string) {
	filename := "l2e.txt"
	data := "Hello There, welcome to Learn2Earn where growth is guarantee!!"
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return
	}
}
