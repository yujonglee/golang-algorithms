package search

import "errors"

func linearSearch(items []int, key int) (int, error) {
	for index, item := range items {
		if item == key {
			return index, nil
		}
	}

	return -1, errors.New("NOT FOUND")
}
