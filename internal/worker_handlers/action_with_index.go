package workerhandlers

import (
	h "edit-tool/pkg/utils"
)

func (w *Worker) ActionWithIndex(input []string) []string {
	h.MergeActionSeparatedWithSpaces(input)
	return nil
}
