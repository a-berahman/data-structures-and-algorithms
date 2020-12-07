package main

import (
	"fmt"

	"github.com/a-berahman/data-structures-and-algorithms/sort"
)

func main() {
	arr := []int{4, 9, 1, 4, 2, 7, 8, 0}
	res := sort.QucikSort(arr)
	fmt.Println(res)
}
