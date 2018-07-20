package main

import (
	"fmt"
	"github.com/shubhamagarwal003/go-algos/sort"
)

func main() {
	arr := []int{5, 2, 1, 2, 3, 10}
	// arr := []int{1, 2, 2, 0, 3, 10}
	sort.MergeSort(arr)
	fmt.Println(arr)
}
