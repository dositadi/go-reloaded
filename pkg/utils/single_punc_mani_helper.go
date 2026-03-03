package utils

import (
	"slices"
	"strings"
)

// This function checks that a Char is a punctuation symbol
func IsPunctuation(input string) bool {
	output := true

	for _, rn := range input {
		if !(rn == ',' || rn == '.' || rn == '?' || rn == ':' || rn == ';' || rn == '!') {
			output = false
			break
		}
	}
	return output
}

// This function checks that a string contains a punctuation in the first index
func ContainsPunctuation(input string) bool {
	output := false

	for i, rn := range input {
		if i == 0 {
			if strings.ContainsAny(string(rn), ",.?:;!") {
				output = true
			}
		}
	}
	return output
}

// This function gets the char at a specific index in the string and deletes that character only if its a punctuation
func GetCharAndDeleteChar(input *string, index int) rune {
	temp := []rune(*input)

	var output rune

	for i, rn := range *input {
		if i == index {
			output = rn
			// Only delete the rune if the output is a rune
			if IsPunctuation(string(output)) {
				temp = slices.Delete(temp, i, i+1)
			}
			break
		}
	}

	*input = string(temp)
	return output
}

// This function appends a character to the end of a string
func AppendChar(input *string, char rune) {
	temp := []rune(*input)

	temp = append(temp, char)

	*input = string(temp)
}

// This function appends multiple characters to the end of a string
func AppendChars(input *string, chars string) {
	temp := *input

	temp += chars

	*input = temp
}
