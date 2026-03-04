package unittest

import (
	w "edit-tool/internal/worker_handlers"
	"slices"
	"strings"
	"testing"
)

func TestActionWithIndex(t *testing.T) {
	worker := w.CreateNewWorker()
	expected := strings.Fields("it (cap) was the best of times, it was the worst of TIMES (low) , it was the age of wisdom, it was the age of FOOLIShness , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter OF DESPAIR . This is SO EXCITING")
	input := strings.Fields("it (cap) was the best of times, it was the worst of TIMES (low) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair (up, 2) . This is so exciting (up, 2)")

	outPut := worker.ActionWithIndex(input)

	if slices.Compare(expected, outPut) != 0 {
		t.Fatalf("The Output\n%+v\nis not the same as\n%+v\nfor Vowel Manipulator!", outPut, expected)
	}
	t.Log("Function is Ok")
}
