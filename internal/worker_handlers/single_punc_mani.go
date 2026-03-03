package workerhandlers

import (
	h "edit-tool/pkg/utils"
	"slices"
)

// Implementation of the SinglePunctuationManipulator func from Engine interface
func (w *Worker) SingleAndGroupOfPunctuationManipulator(input []string) []string {
	output := input

	length := len(input)

	for i := length - 1; i >= 1; i-- {
		currentWord := input[i]

		// Check for a single standalone punctuation
		if len(currentWord) < 2 {
			if h.IsPunctuation(currentWord) {
				punctuation := h.GetCharAndDeleteChar(&output[i], 0)
				h.AppendChar(&output[i-1], punctuation)
				output = slices.Delete(output, i, i+1)
			}
		}

		if len(currentWord) > 2 {
			// Check for group of punctuations
			if h.IsPunctuation(currentWord) {
				h.AppendChars(&output[i-1], currentWord)
				output = slices.Delete(output, i, i+1)
			}
			// Check for words with punctuations in the first index
			if h.ContainsPunctuation(currentWord) {
				if h.ContainsPunctuation(currentWord) {
					punctuation := h.GetCharAndDeleteChar(&output[i], 0)
					if h.IsPunctuation(string(punctuation)) {
						h.AppendChar(&output[i-1], punctuation)
					}
				}
			}
		}
	}
	return output
}
