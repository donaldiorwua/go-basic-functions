package gobasicfunctions

func ManualFields(text string) []string {
	var result []string
	var char []rune

	for _, words := range text {
		if words != ' ' {
			char = append(char, words)
		}else if words == ' ' && len(char) > 0 {
			result = append(result, string(char))
			char = []rune{}
		}
	}
	result = append(result, string(char))
	return result
}
