package config

import (
	m "edit-tool/pkg/models"
	h "edit-tool/pkg/utils"
	"slices"
)

type Engine interface {
	ReadFile(filepath string) ([]string, *m.Error)                  // Function to read the file in the given filepath --> done
	StringBinHexManipulator(input []string) []string                // Function to solve the (cap), (low) and (up) problem --> done
	ActionWithIndex(input []string) []string                        // Function to solve action with index problem such as (cap, 3),(up, 4),(low, 4) -->
	SingleAndGroupOfPunctuationManipulator(input []string) []string // Function to solve the single puntuations problem --> done
	PuntuationMarkManipulator(input []string) []string              // Function to handle the puntuation mark problem --> done
	VowelManipulator(input []string) []string                       // Function to handle the vowel manipulation problem --> done
	WriteResult(output []string, filename string) *m.Error          // Function to write to the output file -->
}

// The EngineWorker struct produce an object (worker) of type Engine (that implements all the methods in the Engine interface)
type EngineWorker struct {
	Worker     Engine
	FilePath   string
	ResultFile string
}

// A constructor to create  new engine worker
func CreateNewEngineWorker(worker Engine, filePaths string, resultFile string) *EngineWorker {
	return &EngineWorker{
		Worker:     worker,
		FilePath:   filePaths,
		ResultFile: resultFile,
	}
}

// The engine function that does the work on a single file
func (e *EngineWorker) Work(filePath string) {
	fileContent, err := e.Worker.ReadFile(filePath)
	if err != nil {
		h.PrintErrorMessage(*err)
		return
	}

	outPut1 := e.DoItTillComplete(fileContent)

	err2 := e.Worker.WriteResult(outPut1, e.ResultFile)
	if err2 != nil {
		h.PrintErrorMessage(*err2)
		return
	}
}

var OutputBuffer []string

func (e *EngineWorker) DoItTillComplete(input []string) []string {
	output1 := e.Worker.ActionWithIndex(input)

	outPut2 := e.Worker.StringBinHexManipulator(output1)

	outPut3 := e.Worker.PuntuationMarkManipulator(outPut2)

	outPut4 := e.Worker.SingleAndGroupOfPunctuationManipulator(outPut3)

	if slices.Compare(outPut4, OutputBuffer) == 0 {
		OutputBuffer = []string{}
		return outPut4
	} else {
		outPut5 := e.Worker.VowelManipulator(outPut4)

		OutputBuffer = outPut5

		return e.DoItTillComplete(outPut5)
	}
}

// Helper function that cordinates the editing work.
func (e *EngineWorker) DoAllWork() {
	e.Work(e.FilePath)
}
