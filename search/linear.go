package search

func linearSearch(items []int, key int) bool {
	for _, item := range items {
		if item == key {
			return true
		}
	}

	return false
}
