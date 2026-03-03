package workerhandlers

import m "edit-tool/pkg/models"

// Implementation of the WriteResult func from Engine interface
func (w *Worker) WriteResult(output []string) *m.Error {
	return nil
}
