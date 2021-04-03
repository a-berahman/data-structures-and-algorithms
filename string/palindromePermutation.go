package string

// IsPalindrome Given a string, write a function to check if it is a permutation of a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation
// is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
// 1.5
// 1.6
// EXAMPLE
// Input: Tact Coa
// Output: True (permutations: "taco cat", "atco eta", etc.)
func IsPalindrome(str string) bool {

	dic := make(map[rune]int)
	odd := 0
	for _, v := range str {
		if v == ' ' {
			continue
		}
		dic[v]++
		if dic[v]%2 != 0 {
			odd++
		} else {
			odd--
		}
	}

	return !(odd > 0)

}
