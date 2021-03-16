package array

// LeetCode
// Given an array, rotate the array to the right by k steps, where k is non-negative.
// ex:
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
func rotate(nums []int, k int) {
	k = k % len(nums)

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}
func reverse(nums []int, i, j int) {
	for i <= j {
		nums[i], nums[j] = nums[j], nums[i]
		j--
		i++
	}
}
 