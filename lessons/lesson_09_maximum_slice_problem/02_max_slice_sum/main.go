package main

// A non-empty array A consisting of N integers is given. A pair of integers (P, Q), such that 0 ≤ P ≤ Q < N, is called a slice of array A. The sum of a slice (P, Q) is the total of A[P] + A[P+1] + ... + A[Q].

// Write a function:

// func Solution(A []int) int

// that, given an array A consisting of N integers, returns the maximum sum of any slice of A.

// For example, given array A such that:

// A[0] = 3  A[1] = 2  A[2] = -6
// A[3] = 4  A[4] = 0
// the function should return 5 because:

// (3, 4) is a slice of A that has sum 4,
// (2, 2) is a slice of A that has sum −6,
// (0, 1) is a slice of A that has sum 5,
// no other slice of A has sum greater than (0, 1).
// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [1..1,000,000];
// each element of array A is an integer within the range [−1,000,000..1,000,000];
// the result will be an integer within the range [−2,147,483,648..2,147,483,647].

import "fmt"

func main() {
	A := []int{3, 2, -6, 4, 0}
	// A := []int{-2, 1}

	result := Solution(A)

	fmt.Println("result:", result)
}

func Solution(A []int) int {
	// write your code in Go 1.4

	if len(A) == 1 {
		return A[0]
	}

	maxSum := A[0]
	partSum := A[0]

	for i := 1; i < len(A); i++ {
		compare := partSum + A[i]
		if A[i] > compare {
			compare = A[i]
		}

		if maxSum < compare {
			maxSum = compare
		}

		if compare < 0 {
			partSum = 0
		} else {
			partSum = compare
		}
	}

	return maxSum
}

// https://app.codility.com/demo/results/trainingDYASA8-P9K/
