package main

import (
	"fmt"
	"sort"
)

// A non-empty array A consisting of N integers is given. The product of triplet (P, Q, R) equates to A[P] * A[Q] * A[R] (0 ≤ P < Q < R < N).

// For example, array A such that:

//   A[0] = -3
//   A[1] = 1
//   A[2] = 2
//   A[3] = -2
//   A[4] = 5
//   A[5] = 6
// contains the following example triplets:

// (0, 1, 2), product is −3 * 1 * 2 = −6
// (1, 2, 4), product is 1 * 2 * 5 = 10
// (2, 4, 5), product is 2 * 5 * 6 = 60
// Your goal is to find the maximal product of any triplet.

// Write a function:

// func Solution(A []int) int

// that, given a non-empty array A, returns the value of the maximal product of any triplet.

// For example, given array A such that:

//   A[0] = -3
//   A[1] = 1
//   A[2] = 2
//   A[3] = -2
//   A[4] = 5
//   A[5] = 6
// the function should return 60, as the product of triplet (2, 4, 5) is maximal.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [3..100,000];
// each element of array A is an integer within the range [−1,000..1,000].

func main() {
	// A := []int{-3, 1, 2, -2, 5, 6}

	A := []int{-5, 5, -5, 4}

	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {

	sort.Ints(A)

	result1 := A[len(A)-3] * A[len(A)-2] * A[len(A)-1]

	if A[0] < 0 && A[1] < 0 && A[len(A)-1] >= 0 {
		result2 := A[0] * A[1] * A[len(A)-1]
		if result1 < result2 {
			return result2
		}
	}

	return result1
}
