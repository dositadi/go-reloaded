package utils

import (
	m "edit-tool/pkg/models"
	"fmt"
	"strconv"
	"strings"
)

func MergeActionSeparatedWithSpaces(input []string) {
	output := input
	length := len(input)

	for i := length - 1; i >= 1; i-- {
		if strings.Contains(input[i], ")") {
			Merge(&output[i-1], output[i-1], output[i])
		}
	}

	fmt.Println(output)
}

func Merge(input *string, first, second string) {
	temp := first + second

	*input = temp
}

// Helper function to convert strings in lower case to upper case
func ToUpperIndices(inputs []string) {
	for i := range inputs {
		ToUpper(&inputs[i])
	}
	fmt.Println(inputs)
}

// Helper function to convert strings to lowercase
func ToLowerIndices(inputs []string) {
	for i := range inputs {
		ToLower(&inputs[i])
	}
	fmt.Println(inputs)
}

// Helper function to convert characters of a string to lowercase
func CapIndices(input *string, indices int) {
	for i := range indices {
		Cap(input, i)
	}
	fmt.Println(input)
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
