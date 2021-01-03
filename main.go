package main

import (
	"fmt"

	"github.com/a-berahman/data-structures-and-algorithms/sort"
)

func main() {
	arr := []int{3, 4, 1, 2, 5}
	res := sort.QucikSortTwo(arr)
	fmt.Println(res)
}
