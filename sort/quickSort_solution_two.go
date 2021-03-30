package sort

import "math/rand"

func QucikSortTwo(arr []int) []int {
	if len(arr) < 2 {
		return nil
	}
	low, high := 0, len(arr)-1
	pivot := rand.Int() % len(arr)
	arr[high], arr[pivot] = arr[pivot], arr[high]
	for i, _ := range arr {
		if arr[i] < arr[high] {
			arr[i], arr[low] = arr[low], arr[i]
			low++
		}
	}
	arr[low], arr[high] = arr[high], arr[low]
	QucikSortTwo(arr[:low])
	QucikSortTwo(arr[low+1:])
	return arr
}
