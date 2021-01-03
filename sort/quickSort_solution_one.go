package sort

func QucikSortOne(arr []int) []int {
	sort(arr, 0, len(arr)-1)
	return arr
}

func sort(arr []int, low, high int) {
	if len(arr) < 2 {
		return
	}
	if low > high {
		return
	}

	index := partition(arr, low, high)
	sort(arr, low, index-1)
	sort(arr, index+1, high)

}

func partition(arr []int, low, high int) int {

	pivot := arr[low]
	for low < high {
		for low < high && arr[high] >= pivot {
			high--
		}
		arr[low], arr[high] = arr[high], arr[low]
		for low < high && arr[low] <= pivot {
			low++
		}
		arr[low], arr[high] = arr[high], arr[low]
	}
	return low

}
