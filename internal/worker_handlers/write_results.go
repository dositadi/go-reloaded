package workerhandlers

import (
	m "edit-tool/pkg/models"
	h "edit-tool/pkg/utils"
	"os"
)

// Implementation of the WriteResult func from Engine interface
func (w *Worker) WriteResult(input []string, filename string) *m.Error {
	path := "output/" + filename
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return &m.Error{
			Err:                  h.SERVER_ERR,
			Details:              h.SERVER_ERR_DETAIL,
			SuggestiveCorrection: h.SERVER_ERR_SUGGESTIVE_CORRECTION,
		}
	}

	for i, word := range input {
		if i != len(input)-1 {
			_, err := file.WriteString(word + " ")
			if err != nil {
				return &m.Error{
					Err:                  h.SERVER_ERR,
					Details:              h.SERVER_ERR_DETAIL,
					SuggestiveCorrection: h.SERVER_ERR_SUGGESTIVE_CORRECTION,
				}
			}
		} else {
			_, err := file.WriteString(word)
			if err != nil {
				return &m.Error{
					Err:                  h.SERVER_ERR,
					Details:              h.SERVER_ERR_DETAIL,
					SuggestiveCorrection: h.SERVER_ERR_SUGGESTIVE_CORRECTION,
				}
			}
		}
	}
	return nil
}
