package utils

import "strings"

// This function takes a slice of strings of strings and converts it to a slice of words.
func StringsToSliceOfWords(input string) []string {
	length := len(input)
	if length == 0 {
		return nil
	}

	output := strings.Fields(input)
	return output
}
