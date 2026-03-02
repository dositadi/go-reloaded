package utils

import (
	"fmt"
	"strings"
	"unicode"
)

func ToUpper(input *string) {
	temp := *input
	temp = strings.ToUpper(temp)
	*input = temp
}

func ToLower(input *string) {
	temp := *input
	temp = strings.ToLower(temp)
	*input = temp
}

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

func ToUpperIndices(inputs []string) {
	for i := range inputs {
		ToUpper(&inputs[i])
	}
	fmt.Println(inputs)
}

func ToLowerIndices(inputs []string) {
	for i := range inputs {
		ToLower(&inputs[i])
	}
	fmt.Println(inputs)
}

func CapIndices(input string, indices int) {
	for i := range indices {
		Cap(&input, i)
	}
	fmt.Println(input)
}
