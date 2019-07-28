package main

// This is a demo task.

// Write a function:

// func Solution(A []int) int

// that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

// For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

// Given A = [1, 2, 3], the function should return 4.

// Given A = [−1, −3], the function should return 1.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [1..100,000];
// each element of array A is an integer within the range [−1,000,000..1,000,000].

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
