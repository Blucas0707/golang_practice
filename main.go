package main

import (
	"fmt"
	"golang_practice/sort/selectionsort"
)

func main() {
	nums := []int{7, 14, 2, 3, 123, 4, 324, 234, 3, 31, 0, 4, 324, 4, 546, 56, 78, 78, 7, 5, 3}
	fmt.Println(selectionsort.Selectionsort(nums))
	// fmt.Println(howsum.MemoHowSum(300, []int{7, 14, 286}, make(map[int][]int)))
}
