package string

// Given a string, find the first non-repeating character in it and return its index. If it doesn't exist, return -1.
// ex:
// 	s = "leetcode"
// 	return 0.
//
// 	s = "loveleetcode"
// 	return 2.
func FirstUniqChar(s string) int {
	dic := make(map[rune]int)
	for _, v := range s {
		if _, ok := dic[v]; ok {
			dic[v] += 1
		} else {
			dic[v] = 1
		}
	}

	for j, v := range s {
		if dic[v] == 1 {
			return j
		}

	}

	return 0
}
