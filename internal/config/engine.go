package config

import (
	m "edit-tool/pkg/models"
	h "edit-tool/pkg/utils"
)

type Engine interface {
	ReadFile(filepath string) ([]string, *m.Error)                  // Function to read the file in the given filepath --> done
	StringBinHexManipulator(input []string) []string                // Function to solve the (cap), (low) and (up) problem --> done
	ActionWithIndex(input []string) []string                        // Function to solve action with index problem such as (cap, 3),(up, 4),(low, 4) -->
	SingleAndGroupOfPunctuationManipulator(input []string) []string // Function to solve the single puntuations problem --> done
	PuntuationMarkManipulator(input []string) []string              // Function to handle the puntuation mark problem --> done
	VowelManipulator(input []string) []string                       // Function to handle the vowel manipulation problem --> done
	WriteResult(output []string) *m.Error                           // Function to write to the output file -->
}

// The EngineWorker struct produce an object (worker) of type Engine (that implements all the methods in the Engine interface)
type EngineWorker struct {
	Worker    Engine
	FilePaths []string
}

// A constructor to create  new engine worker
func CreateNewEngineWorker(worker Engine, filePaths []string) *EngineWorker {
	return &EngineWorker{
		Worker:    worker,
		FilePaths: filePaths,
	}
}

// The engine function that does the work on a single file
func (e *EngineWorker) Work(filePath string) {
	fileContent, err := e.Worker.ReadFile(filePath)
	if err != nil {
		h.PrintErrorMessage(*err)
		return
	}

	output1 := e.Worker.ActionWithIndex(fileContent)

	outPut2 := e.Worker.StringBinHexManipulator(output1)

	outPut3 := e.Worker.PuntuationMarkManipulator(outPut2)

	outPut4 := e.Worker.SingleAndGroupOfPunctuationManipulator(outPut3)

	outPut5 := e.Worker.VowelManipulator(outPut4)

	err2 := e.Worker.WriteResult(outPut5)
	if err2 != nil {
		h.PrintErrorMessage(*err2)
		return
	}
}

// Helper function that cordinates the editing work.
func (e *EngineWorker) DoAllWork() {
	for _, filepath := range e.FilePaths {
		e.Work(filepath)
	}
}
