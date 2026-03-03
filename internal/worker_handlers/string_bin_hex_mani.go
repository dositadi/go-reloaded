package workerhandlers

import (
	h "edit-tool/pkg/utils"
	"slices"
)

// Implementation of the StringManipulator func from Engine interface
func (w *Worker) StringBinHexManipulator(input []string) []string {
	output := input
	length := len(output)

	for i := 0; i < length; i++ {
		currentWord := input[i]

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
		case "(bin)":
			h.Bin(&output[i-1])
			output = slices.Delete(output, i, i+1)
		case "(hex)":
			h.Hex(&output[i-1])
			output = slices.Delete(output, i, i+1)
		}
	}
	return output
}
