package workerhandlers

import (
	m "edit-tool/pkg/models"
	h "edit-tool/pkg/utils"
	"os"
)

// Implementation of the ReadFile func from Engine
func (w *Worker) ReadFile(filepath string) ([]string, *m.Error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644) // Function to open and read from file.
	if err != nil {
		return nil, &m.Error{ // If error not equal nil, return the custom error struct to the user.
			Err:                  h.FILE_READ_ERR,
			Details:              h.FILE_READ_ERR_DETAIL,
			SuggestiveCorrection: h.ARG_ERR_SUGGESTIVE_CORRECTION,
		}
	}

	fileStat, err2 := file.Stat() // Get the file stat information
	if err2 != nil {
		return nil, &m.Error{
			Err:                  h.FILE_READ_ERR,
			Details:              h.FILE_READ_ERR_DETAIL,
			SuggestiveCorrection: h.FILE_READ_ERR_SUGGESTIVE_CORRECTION,
		}
	}

	fileBuffer := make([]byte, fileStat.Size()) // outputoral buffer to store the contents of the file

	_, err3 := file.Read(fileBuffer) // Read the file content and store in the fileBuffer
	if err3 != nil {
		return nil, &m.Error{
			Err:                  h.FILE_READ_ERR,
			Details:              h.FILE_READ_ERR_DETAIL,
			SuggestiveCorrection: h.FILE_READ_ERR_SUGGESTIVE_CORRECTION,
		}
	}

	if len(fileBuffer) == 0 { // Check to ensure that the file is not empty
		return nil, &m.Error{
			Err:                  h.EMPTY_FILE_ERR,
			Details:              h.EMPTY_FILE_ERR_DETAIL,
			SuggestiveCorrection: h.EMPTY_FILE_ERR_SUGGESTIVE_CORRECTION,
		}
	}
	result := h.StringsToSliceOfWords(string(fileBuffer)) // Convert the file content into a slice of words
	return result, nil
}
