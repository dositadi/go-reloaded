package unittest

import (
	w "edit-tool/internal/worker_handlers"
	"slices"
	"strings"
	"testing"
)

func TestSingleAndGroupPunct(t *testing.T) {
	worker := w.CreateNewWorker()
	expected := strings.Fields("Punctuation tests are... kinda boring, what do you think? I was thinking... You were right, I was sitting over there, and then BAMM!!")
	input := strings.Fields("Punctuation tests are ... kinda boring ,what do you think ?  I was thinking ... You were right , I was sitting over there ,and then BAMM !! ")

	outPut := worker.SingleAndGroupOfPunctuationManipulator(input)

	if slices.Compare(expected, outPut) != 0 {
		t.Fatalf("The Output\n%+v\nis not the same as\n%+v\nfor Vowel Manipulator!", outPut, expected)
	}
	t.Log("Function is Ok")
}
