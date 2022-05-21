package search

import "testing"

func TestBinaryLoopFound(t *testing.T) {
	index, error := binaryLoopSearch([]int{1, 2, 3, 5, 8, 11}, 8)

	if index != 4 || error != nil {
		t.Error()
	}
}

func TestBinaryLoopNotFound(t *testing.T) {
	index, error := binaryLoopSearch([]int{1, 2, 3, 5, 8, 11}, 4)

	if index != -1 || error == nil {
		t.Error()
	}
}

func TestBinaryRecursiveFound(t *testing.T) {
	index, error := binaryRecursiveSearch([]int{1, 2, 3, 5, 8, 11}, 8, 0, 5)

	if index != 4 || error != nil {
		t.Error()
	}
}

func TestBinaryRecursiveNotFound(t *testing.T) {
	index, error := binaryRecursiveSearch([]int{1, 2, 3, 5, 8, 11}, 4, 0, 5)

	if index != -1 || error == nil {
		t.Error()
	}
}
