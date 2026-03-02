package config

import (
	w "edit-tool/internal/worker_handlers"
	m "edit-tool/pkg/models"
	h "edit-tool/pkg/utils"
	"os"
)

type App struct{}

// Function to collect user filePath strings
func (a *App) GetFileNames() ([]string, *m.Error) {
	if len(os.Args) < 2 {
		return nil, &m.Error{
			Err:                  h.ARG_ERR,
			Details:              h.ARG_ERR_DETAIL,
			SuggestiveCorrection: h.ARG_ERR_SUGGESTIVE_CORRECTION,
		}
	}

	args := os.Args[1:]

	return args, nil
}

// The start button to kick start engine
func (a *App) Run() {
	filePaths, err := a.GetFileNames()
	if err != nil {
		h.PrintErrorMessage(*err)
		return
	}

	worker := w.CreateNewWorker()
	engineWorker := CreateNewEngineWorker(worker, filePaths)
	engineWorker.DoAllWork()
}
