package search

import "testing"

func Test(t *testing.T) {
	if !linearSearch([]int{1, 2, 3}, 2) {
		t.Error()
	}

	if linearSearch([]int{1, 2, 3}, 4) {
		t.Error()
	}
}
