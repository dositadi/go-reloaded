package utils

import (
	m "edit-tool/pkg/models"
	"slices"
	"strconv"
	"strings"
)

// This function cleans the input array for use.
func MergeActionSeparatedWithSpaces(input []string) []string {
	output := input
	length := len(input)

	for i := length - 1; i >= 1; i-- {
		if CheckChar(input[i], ')') && CheckChar(input[i-1], '(') {
			Merge(&output[i-1], output[i-1], output[i])
			output = slices.Delete(output, i, i+1)
		}
	}

	return output
}

// This function checks a string if it has a specific char
func CheckChar(input string, char rune) bool {
	temp := []rune(input)

	for _, rn := range temp {
		if rn == char {
			return true
		}
	}
	return false
}

// This function merges two strings and stores the output in input
func Merge(input *string, first, second string) {
	temp := first + second

	*input = temp
}

// Helper function to convert characters of a string to lowercase
func CapIndices(input *string, indices int) {
	temp := *input
	for i := range indices {
		Cap(&temp, i)
	}
	*input = temp
}

// Function to check for strings that follow the pattern (low,2), (up,6),(cap,4)
func ContainsSpecialChar(input string) bool {
	if strings.HasPrefix(input, "(") && strings.HasSuffix(input, ")") && strings.ContainsRune(input, ',') {
		return true
	}
	return false
}

// This function splits the string and gets the action and indices
func SplitAndGetIndices(input string) (string, int, *m.Error) {
	valWithoutPrefix, _ := strings.CutPrefix(input, "(")
	valueWithoutPrefixAndSuffix, _ := strings.CutSuffix(valWithoutPrefix, ")")
	cleanedValue := strings.Split(valueWithoutPrefixAndSuffix, ",")

	if len(cleanedValue) != 2 {
		return "", 0, &m.Error{
			Err:                  INVALID_SYNTAX_ERR,
			Details:              INVALID_SYNTAX_ERR_DETAIL,
			SuggestiveCorrection: INVALID_SYNTAX_ERR_SUGGESTIVE_CORRECTION,
		}
	}

	var (
		indiceStr = ""
		action    = ""
	)

	for i := range cleanedValue {
		if i == 0 {
			indiceStr = cleanedValue[1]
		} else {
			action = cleanedValue[0]
		}
	}

	indices, err := strconv.Atoi(indiceStr)
	if err != nil {
		return "", 0, &m.Error{
			Err:                  INVALID_INDICES_ERROR,
			Details:              INVALID_INDICES_ERROR_DETAIL,
			SuggestiveCorrection: INVALID_INDICES_ERROR_SUGGESTIVE_CORRECTION,
		}
	}

	return action, indices, nil
}
