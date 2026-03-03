package workerhandlers

import (
	h "edit-tool/pkg/utils"
	"fmt"
)

// Implementation of the PuntuationMarkManipulator func from Engine interface
func (w *Worker) PuntuationMarkManipulator(input []string) []string {
	fmt.Println(input)

	indices := h.NotePMIndices(input)

	h.FixPunctMarks(input, indices)
	return nil
}
