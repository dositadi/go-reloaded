package workerhandlers

import (
	h "edit-tool/pkg/utils"
	"fmt"
	"strings"
	"unicode"
)

// Implementation of the VowelManipulatior func from Engine interface
func (w *Worker) VowelManipulator(input []string) []string {
	output := input
	lenght := len(input)

	for i := lenght - 1; i >= 1; i-- {
		currentWord := input[i]

		// Coverting the charcater to lowercase for uniformity in check
		char := unicode.ToLower(h.GetFirstCharOfString(currentWord))

		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'h' {
			if len(input[i-1]) == 1 && strings.ToLower(input[i-1]) == "a" {
				h.ReplacePreviousWord(&output[i-1])
			}
		}
	}
	fmt.Println(output)
	return output
}
