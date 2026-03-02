package workerhandlers

import (
	m "edit-tool/pkg/models"
	h "edit-tool/pkg/utils"
	"fmt"
	"os"
	"slices"
)

type Worker struct{}

// This constructor creates a new worker
func CreateNewWorker() *Worker {
	return &Worker{}
}

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
	fmt.Println(result, " ", len(result))
	return result, nil
}

// Implementation of the StringManipulator func from Engine interface
func (w *Worker) StringManipulator(input []string) []string {
	output := input
	//var output []string
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
		}
	}

	//output = output
	fmt.Println(output, " ", len(output))
	return output
}

// Implementation of the BinHexManipulator func from Engine interface
func (w *Worker) BinHexManipulator(input []string) []string {
	return nil
}

// Implementation of the SinglePunctuationManipulator func from Engine interface
func (w *Worker) SinglePunctuationManipulator(input []string) []string {
	return nil
}

// Implementation of the GroupOfPunctuationManipulator func from Engine interface
func (w *Worker) GroupOfPunctuationManipulator(input []string) []string {
	return nil
}

// Implementation of the PuntuationMarkManipulator func from Engine interface
func (w *Worker) PuntuationMarkManipulator(input []string) []string {
	return nil
}

// Implementation of the VowelManipulatior func from Engine interface
func (w *Worker) VowelManipulator(input []string) []string {
	return nil
}

// Implementation of the WriteResult func from Engine interface
func (w *Worker) WriteResult(output []string) *m.Error {
	return nil
}
