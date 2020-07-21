package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := make([]int, 4)
	arr[0] = 12
	arr[1] = 400
	arr[2] = 30
	arr[3] = 2
	fmt.Println("Initial slice:", arr)

	sort.Ints(arr)
	fmt.Println("After sorting (asc):", arr)

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println("After sorting (dsc):", arr)
}
