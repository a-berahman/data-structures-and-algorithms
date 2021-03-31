package string

// Write a method to replace all spaces in a string with '%20'. You may assume that the string
// has sufficient space at the end to hold the additional characters, and that you are given the "true"
// length of the string. (Note: If implementing in Java, please use a character array so that you can
// perform this operation in place.)
func URLify(str string, n int) string {

	space := 0
	for _, s := range str {
		if s == ' ' {
			space++
		}
	}

	index := n + (space * 2)
	arr := make([]rune, n+(space*2))

	for i := n - 1; i >= 0; i-- {
		if str[i] == ' ' {
			arr[index-1] = '0'
			arr[index-2] = '2'
			arr[index-3] = '%'
			index = index - 3
		} else {
			arr[index-1] = rune(str[i])
			index--
		}
	}

	return string(arr)
}
