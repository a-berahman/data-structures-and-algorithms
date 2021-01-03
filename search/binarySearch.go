package search

import "fmt"

func binarySearch(arr []int, item int) (int, error) {
	if len(arr) < 1 {
		return 0, fmt.Errorf("item not found")
	}
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high - low) / 2)
		val := arr[mid]
		if val == item {
			return mid, nil
		}
		if val < item {
			low = mid + 1
		}
		if val > item {
			high = mid - 1
		}
	}
	return 0, fmt.Errorf("item not found")
}
