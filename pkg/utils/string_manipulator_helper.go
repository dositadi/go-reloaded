package utils

import (
	m "edit-tool/pkg/models"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Helper function to convert string to upper case
func ToUpper(input *string) {
	temp := *input
	temp = strings.ToUpper(temp)
	*input = temp
}

// Helper function to convert string to lowercase
func ToLower(input *string) {
	temp := *input
	temp = strings.ToLower(temp)
	*input = temp
}

// Helper function to convert specific characters of a string to upper case
func Cap(input *string, index int) {
	var temp []rune
	for n, char := range *input {
		if n == index {
			char = unicode.ToUpper(char)
			temp = append(temp, char)
		} else {
			temp = append(temp, char)
		}
	}
	*input = string(temp)
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
