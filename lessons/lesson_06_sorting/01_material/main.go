package main

// You are given a zero-indexed array A consisting of n > 0 integers; you must return
// the number of unique values in array A.

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{1, 4, 2, 3, 2, 5, 4}

	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {

	sort.Ints(A)
	result := 1
	for i := 1; i < len(A); i++ {
		if A[i] != A[i-1] {
			result++
		}
	}
	return result
}
