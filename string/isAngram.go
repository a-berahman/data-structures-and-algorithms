package string

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// Example 1:
//
// Input: s = "anagram", t = "nagaram"
// Output: true
// Example 2:
//
// Input: s = "rat", t = "car"
// Output: false
func isAnagram_First_Solution(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	first := make(map[rune]int)
	second := make(map[rune]int)
	for _, v := range s {
		if _, ok := first[v]; ok {
			first[v] += 1
		} else {
			first[v] = 1
		}
	}
	for _, v := range t {
		if _, ok := second[v]; ok {
			second[v] += 1
		} else {
			second[v] = 1
		}
	}

	for key, value := range first {
		if _, ok := second[key]; !ok {
			return false
		}
		if second[key] != value {
			return false
		}
	}
	return true

}
func isAnagram_Second_Solution(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dic := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		dic[s[i]]++
		dic[t[i]]--

	}
	for _, v := range dic {
		if v != 0 {
			return false
		}
	}
	return true

}
