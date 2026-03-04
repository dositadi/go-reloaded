package unittest

import (
	w "edit-tool/internal/worker_handlers"
	"slices"
	"strings"
	"testing"
)

func TestPunctMark(t *testing.T) {
	worker := w.CreateNewWorker()
	expected := strings.Fields("I am 'exactly' 'how they describe me: awesome' . As Elton John said: 'I am the most well-known homosexual in the world'")
	input := strings.Fields("I am ' exactly ' ' how they describe me: awesome ' . As Elton John said: ' I am the most well-known homosexual in the world '")

	outPut := worker.PuntuationMarkManipulator(input)

	if slices.Compare(expected, outPut) != 0 {
		t.Fatalf("The Output\n%+v\nis not the same as\n%+v\nfor Vowel Manipulator!", outPut, expected)
	}
	t.Log("Function is Ok")
}
