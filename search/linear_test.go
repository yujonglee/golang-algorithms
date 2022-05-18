package search

import "testing"

func TestLinearFound(t *testing.T) {
	index, error := linearSearch([]int{1, 2, 3}, 2)

	if index != 1 || error != nil {
		t.Error()
	}
}

func TestLinearNotFound(t *testing.T) {
	index, error := linearSearch([]int{1, 2, 3}, 4)

	if index != -1 || error == nil {
		t.Error()
	}
}
