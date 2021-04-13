package string

import (
	"strconv"
	"strings"
)

// StringCompression: Implement a method to perform basic string compression using the counts
// of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the
// "compressed" string would not become smaller than the original string, your method should return
// the original string. You can assume the string has only uppercase and lowercase letters (a - z).
func StringCompression(str string) string {
	builder := strings.Builder{}
	counter := 0
	for i := 0; i < len(str); i++ {
		counter++
		if i+1 >= len(str) {
			builder.WriteString(string(str[i]))
			builder.WriteString(strconv.Itoa(counter))
			counter = 0
			break
		} else if str[i] != str[i+1] {

			builder.WriteString(string(str[i]))
			builder.WriteString(strconv.Itoa(counter))
			counter = 0

		}
	}
	if len(builder.String()) < len(str) {
		return builder.String()
	}
	return str
}
