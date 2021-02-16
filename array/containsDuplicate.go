package array

// LeetCode
// Given an array of integers, find if the array contains any duplicates.
// Your function should return true if any value appears at least twice in the array,
// and it should return false if every element is distinct.
func containsDuplicate(nums []int) bool {
	dic := make(map[int]bool)
	for _, v := range nums {
		if _, ok := dic[v]; ok {
			return true
		}
		dic[v] = true
	}
	return false

}
