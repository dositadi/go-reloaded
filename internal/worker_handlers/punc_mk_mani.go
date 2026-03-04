package workerhandlers

import (
	h "edit-tool/pkg/utils"
)

// Implementation of the PuntuationMarkManipulator func from Engine interface
func (w *Worker) PuntuationMarkManipulator(input []string) []string {
	indices := h.NotePMIndices(input)

	output := h.FixPunctMarks(input, indices)
	return output
}
