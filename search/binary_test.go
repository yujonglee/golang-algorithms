package search

import "testing"

func TestBinaryFound(t *testing.T) {
	index, error := binaryLoopSearch([]int{1, 2, 3, 5, 8, 11}, 8)

	if index != 4 || error != nil {
		t.Error()
	}
}

func TestBinaryNotFound(t *testing.T) {
	index, error := binaryLoopSearch([]int{1, 2, 3, 5, 8, 11}, 4)

	if index != -1 || error == nil {
		t.Error()
	}
}
