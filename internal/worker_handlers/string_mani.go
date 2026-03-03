package workerhandlers

import (
	h "edit-tool/pkg/utils"
	"fmt"
	"slices"
	"strings"
)

// Implementation of the StringManipulator func from Engine interface
func (w *Worker) StringManipulator(input []string) []string {
	output := input
	length := len(output)

	for i := 0; i < length; i++ {
		currentWord := input[i]
		//fmt.Println(currentWord)

		if i < len(input)-1 && strings.HasPrefix(input[i], "(") {
			if i < len(input)-2 && strings.HasSuffix(input[i+1], ")") {
				currentWord = input[i] + input[i+1]
				output[i] = currentWord
				output = slices.Delete(output, i+1, i+2)
				fmt.Println(currentWord)

				if h.ContainsSpecialChar(currentWord) {
					action, indices, err := h.SplitAndGetIndices(currentWord)
					fmt.Println(action, indices)
					if err != nil {
						h.PrintErrorMessage(*err)
						return nil
					}

					switch action {
					case "up":
						for j := 1; j <= indices; i++ {
							h.ToUpper(&output[i-j])
						}
						output = slices.Delete(output, i+1, i+1)
					case "low":
						for j := 1; j <= indices; i++ {
							h.ToLower(&output[i-j])
						}
						output = slices.Delete(output, i, i+1)
					case "cap":
						h.CapIndices(&output[i-1], indices)
						fmt.Println(output[i-1])
						output = slices.Delete(output, i, i+1)
					}
				}
			}
		} else {
			switch currentWord {
			case "(up)":
				h.ToUpper(&output[i-1])
				output = slices.Delete(output, i, i+1)
			case "(low)":
				h.ToLower(&output[i-1])
				output = slices.Delete(output, i, i+1)
			case "(cap)":
				h.Cap(&output[i-1], 0)
				output = slices.Delete(output, i, i+1)
			}
		}
	}

	//output = output
	fmt.Println(output, " ", len(output))
	return output
}
