package unittest

import (
	w "edit-tool/internal/worker_handlers"
	"slices"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	worker := w.CreateNewWorker()
	filePath := "/home/gamp/L2E Fellowship/piscine-prompt/go-reloaded/go-reloaded/asset/sampleFile1.txt"
	expected := strings.Fields("it (cap) was the best of times, it was the worst of TIMES (low) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair (up, 2) . This is so exciting (up, 2)")

	outPut, _ := worker.ReadFile(filePath)

	if slices.Compare(expected, outPut) != 0 {
		t.Fatalf("The Output\n%+v\nis not the same as\n%+v\nfor Vowel Manipulator!", outPut, expected)
	}
	t.Log("Function is Ok")
}
