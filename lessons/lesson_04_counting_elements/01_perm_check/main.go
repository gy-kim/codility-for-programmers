package main

import (
	"fmt"
	"sort"
)

// A non-empty array A consisting of N integers is given.

// A permutation is a sequence containing each element from 1 to N once, and only once.

// For example, array A such that:

//     A[0] = 4
//     A[1] = 1
//     A[2] = 3
//     A[3] = 2
// is a permutation, but array A such that:

//     A[0] = 4
//     A[1] = 1
//     A[2] = 3
// is not a permutation, because value 2 is missing.

// The goal is to check whether array A is a permutation.

// Write a function:

// class Solution { public int solution(int[] A); }

// that, given an array A, returns 1 if array A is a permutation and 0 if it is not.

// For example, given array A such that:

//     A[0] = 4
//     A[1] = 1
//     A[2] = 3
//     A[3] = 2
// the function should return 1.

// Given array A such that:

//     A[0] = 4
//     A[1] = 1
//     A[2] = 3
// the function should return 0.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [1..100,000];
// each element of array A is an integer within the range [1..1,000,000,000].

func main() {
	A := []int{4, 2, 3}

	fmt.Println(A)
	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {

	if len(A) == 0 {
		return 0
	}

	sort.Ints(A)

	if A[0] != 1 {
		return 0
	}

	gap := A[len(A)-1] - A[0]
	if gap != len(A)-1 {
		return 0
	}

	for i := 0; i < len(A)-1; i++ {
		if (A[i+1] - A[i]) != 1 {
			return 0
		}
	}

	return 1

}
