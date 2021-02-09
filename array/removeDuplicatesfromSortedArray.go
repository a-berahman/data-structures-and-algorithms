package array

//RemoveDuplicates ...
func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
