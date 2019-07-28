package main

import (
	"fmt"
	"sort"
)

func main() {
	// A := []int{1, 3, 6, 4, 1, 2}
	// A := []int{1, 2, 3}
	A := []int{-1, -3}

	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {

	sort.Ints(A)

	result := 1

	if A[len(A)-1] < 0 {
		return result
	}

	for _, v := range A {
		if v < 0 {
			continue
		}

		if result < v {
			return result
		}

		result = v
		result++

	}

	return result
}
