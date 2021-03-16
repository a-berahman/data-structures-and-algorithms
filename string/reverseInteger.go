package string

import "strconv"

// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	num := x
	if num < 0 {
		num *= -1
	}

	res := ""
	for num/10 != 0 {
		res += strconv.Itoa(num % 10)
		num = num / 10
	}

	res += strconv.Itoa(num % 10)

	digit, _ := strconv.Atoi(res)

	if digit >= 2147483647 {
		return 0
	}
	if x < 0 {
		digit *= -1
	}

	return digit

}
