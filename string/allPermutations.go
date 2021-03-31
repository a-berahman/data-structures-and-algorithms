package string

import "fmt"

func GenerateAllPermutations(val []rune, left, right int) {
	if left == right {
		fmt.Println(string(val))
	} else {
		for i := left; i <= right; i++ {
			val[left], val[i] = val[i], val[left]
			GenerateAllPermutations(val, left+1, right)
			val[left], val[i] = val[i], val[left]
		}
	}

}
