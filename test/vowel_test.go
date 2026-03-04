package unittest

import (
	w "edit-tool/internal/worker_handlers"
	"slices"
	"strings"
	"testing"
)

func TestVowel(t *testing.T) {
	worker := w.CreateNewWorker()
	expected := strings.Fields("There is no greater agony than bearing an untold story inside you. There it was. An amazing rock! An apple was sold for $10.")
	input := strings.Fields("There is no greater agony than bearing a untold story inside you. There it was. A amazing rock! A apple was sold for $10.")

	outPut := worker.VowelManipulator(input)

	if slices.Compare(expected, outPut) != 0 {
		t.Fatalf("The Output\n%+v\nis not the same as\n%+v\nfor Vowel Manipulator!", outPut, expected)
	}
	t.Log("Function is Ok")
}
