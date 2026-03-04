package workerhandlers

import (
	h "edit-tool/pkg/utils"
	"slices"
)

func (w *Worker) ActionWithIndex(input []string) []string {
	corrected := h.MergeActionSeparatedWithSpaces(input)
	lenght := len(corrected)
	output := corrected

	for i := lenght - 1; i >= 1; i-- {
		currentWord := corrected[i]
		if h.ContainsSpecialChar(currentWord) {
			action, num, err := h.SplitAndGetIndices(currentWord)
			if err != nil {
				h.PrintErrorMessage(*err)
				return nil
			}

			switch action {
			case "up":
				start := i - num

				for j := 0; j < len(corrected); j++ {
					if j >= start && j < i {
						h.ToUpper(&output[j])
					}
					if j == i {
						break
					}
				}
				output = slices.Delete(output, i, i+1)
			case "low":
				start := i - num

				for j := 0; j < len(corrected); j++ {
					if j >= start && j < i {
						h.ToLower(&output[j])
					}
					if j == i {
						break
					}
				}
				output = slices.Delete(output, i, i+1)
			case "cap":
				start := i - num

				for j := 0; j < len(corrected); j++ {
					if j >= start && j < i {
						h.Cap(&output[j], 0)
					}
					if j == i {
						break
					}
				}
				output = slices.Delete(output, i, i+1)
			}
		}
	}
	return output
}
