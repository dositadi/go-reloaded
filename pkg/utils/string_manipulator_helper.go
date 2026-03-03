package utils

import (
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

// This function converts the input to its binary equivalent
func Bin(input *string) {
	temp := *input

	val, _ := strconv.ParseInt(temp, 2, 0)

	bin := strconv.Itoa(int(val))

	*input = bin
}

// This function converts the input to its hexadecimal equivalent
func Hex(input *string) {
	temp := *input

	val, _ := strconv.ParseInt(temp, 16, 0)

	hex := strconv.Itoa(int(val))

	*input = hex
}
