package array

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// Follow up: Could you implement a solution with a linear runtime complexity and without using extra memory?
func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums {
		a ^= v
	}
	return a
}
