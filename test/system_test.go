package unittest

import (
	w "edit-tool/internal/worker_handlers"
	"slices"
	"strings"
	"testing"
)

func TestSystem(t *testing.T) {
	worker := w.CreateNewWorker()
	inputs := []string{
		"/home/gamp/L2E Fellowship/piscine-prompt/go-reloaded/go-reloaded/asset/sampleFile1.txt",
		"/home/gamp/L2E Fellowship/piscine-prompt/go-reloaded/go-reloaded/asset/sampleFile2.txt",
		"/home/gamp/L2E Fellowship/piscine-prompt/go-reloaded/go-reloaded/asset/sampleFile4.txt",
		"/home/gamp/L2E Fellowship/piscine-prompt/go-reloaded/go-reloaded/asset/sampleFile5.txt",
		"/home/gamp/L2E Fellowship/piscine-prompt/go-reloaded/go-reloaded/asset/sampleFile6.txt",
	}

	expected := [][]string{
		strings.Fields("It was the best of times, it was the worst of times, it was the age of wisdom, it was the age of FOOLIShness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter OF DESPAIR. This is SO EXCITING."),
		strings.Fields("Simply add 66 and 2 and you will see the result is 68."),
		strings.Fields("There is no greater agony than bearing an untold story inside you. There it was. An amazing rock! An apple was sold for $10."),
		strings.Fields("Punctuation tests are... kinda boring, what do you think? I was thinking... You were right, I was sitting over there, and then BAMM!!"),
		strings.Fields("I am 'exactly' 'how they describe me: awesome'. As Elton John said: 'I am the most well-known homosexual in the world'"),
	}

	var outPut []string

	for i, path := range inputs {
		filecontent, _ := worker.ReadFile(path)
		output1 := worker.ActionWithIndex(filecontent)
		output2 := worker.StringBinHexManipulator(output1)
		output3 := worker.PuntuationMarkManipulator(output2)
		output4 := worker.SingleAndGroupOfPunctuationManipulator(output3)
		outPut = worker.VowelManipulator(output4)

		if slices.Compare(outPut, expected[i]) != 0 {
			t.Errorf("The output of %v\n%+v\ndoes not equal\n%+v", path, outPut, expected[i])
		}
	}

}
