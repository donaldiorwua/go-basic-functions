package gobasicfunctions

func ManualFields(text string) []string {
	var result, char []string

	for _, words := range text {
		if words != ' ' {
			char = append(char, string(words))
		} else if words == ' ' && len(string(words)) > 0 {
			result = append(result, string(words))
			//words = []rune{}
			continue
		} else {
			result = append(result, string(words))
		}
	}
	return result
}
