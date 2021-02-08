package array

//LengthOfLongestSubstring ...
func LengthOfLongestSubstring(str string) int {
	dic := make(map[rune]int)
	s := []rune(str)
	max := 0
	for i, j := 0, 0; j < len(s); j++ {
		if _, ok := dic[s[j]]; ok {
			i = maximum(i, dic[s[j]])
		}
		max = maximum(max, j-i+1)
		dic[s[j]] = j + 1
	}
	return max
}

func maximum(x, y int) int {
	if x > y {
		return x
	}
	return y
}
