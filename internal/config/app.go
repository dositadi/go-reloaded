package config

import (
	w "edit-tool/internal/worker_handlers"
	m "edit-tool/pkg/models"
	h "edit-tool/pkg/utils"
	"os"
)

type App struct{}

// Function to collect user filePath strings
func (a *App) GetFileNames() (string, string, *m.Error) {
	if len(os.Args) < 2 {
		return "", "", &m.Error{
			Err:                  h.ARG_ERR,
			Details:              h.ARG_ERR_DETAIL,
			SuggestiveCorrection: h.ARG_ERR_SUGGESTIVE_CORRECTION,
		}
	}

	args := os.Args[1:]

	if len(args) > 2 {
		return "", "", &m.Error{
			Err:                  h.ARG_ERR2,
			Details:              h.ARG_ERR_DETAIL2,
			SuggestiveCorrection: h.ARG_ERR_SUGGESTIVE_CORRECTION2,
		}
	}

	var filePath, resultFile string

	for i, path := range args {
		if i == 0 {
			filePath = path
		} else {
			resultFile = path
		}
	}

	return filePath, resultFile, nil
}

// The start button to kick start engine
func (a *App) Run() {
	filePath, resultFile, err := a.GetFileNames()
	if err != nil {
		h.PrintErrorMessage(*err)
		return
	}

	worker := w.CreateNewWorker()
	engineWorker := CreateNewEngineWorker(worker, filePath, resultFile)
	engineWorker.DoAllWork()
}
