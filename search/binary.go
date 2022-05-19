package search

import "errors"

func binaryLoopSearch(items []int, key int) (int, error) {
	start := 0
	end := len(items) - 1

	for mid := (start + end) / 2; start <= end; mid = (start + end) / 2 {

		if items[mid] > key {
			end = mid - 1
			continue
		}

		if items[mid] < key {
			start = mid + 1
			continue
		}

		return mid, nil
	}

	return -1, errors.New("NOT FOUND")
}
