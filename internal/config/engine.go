package config

import (
	m "edit-tool/pkg/models"
	h "edit-tool/pkg/utils"
)

type Engine interface {
	ReadFile(filepath string) ([]string, *m.Error)         // Function to read the file in the given filepath
	StringManipulator(input []string) []string             // Function to solve the (cap), (low) and (up) problem
	BinHexManipulator(input []string) []string             // Function to solve the (bin) and (hex) problem
	SinglePunctuationManipulator(input []string) []string  // Function to solve the single puntuations problem
	GroupOfPunctuationManipulator(input []string) []string // Function to solve the group of puntuations problem
	PuntuationMarkManipulator(input []string) []string     // Function to handle the puntuation mark problem
	VowelManipulator(input []string) []string              // Function to handle the vowel manipulation problem
	WriteResult(output []string) *m.Error                  // Function to write to the output file
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

	outPut1 := e.Worker.StringManipulator(fileContent)

	outPut2 := e.Worker.BinHexManipulator(outPut1)

	outPut3 := e.Worker.SinglePunctuationManipulator(outPut2)

	outPut4 := e.Worker.GroupOfPunctuationManipulator(outPut3)

	outPut5 := e.Worker.PuntuationMarkManipulator(outPut4)

	outPut6 := e.Worker.VowelManipulator(outPut5)

	err2 := e.Worker.WriteResult(outPut6)
	if err2 != nil {
		h.PrintErrorMessage(*err2)
		return
	}
}

//
func (e *EngineWorker) DoAllWork() {
	for _, filepath := range e.FilePaths {
		e.Work(filepath)
	}
}
