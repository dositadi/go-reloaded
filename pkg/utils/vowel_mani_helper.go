package utils

func GetFirstCharOfString(input string) rune {
	var output rune
	for i, rn := range input {
		if i == 0 {
			output = rn
			break
		}
	}
	return output
}

func ReplacePreviousWord(input *string) {
	temp := *input

	temp += "n"

	*input = temp
}
