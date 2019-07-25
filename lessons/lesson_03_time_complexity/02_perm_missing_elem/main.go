package main

import "fmt"

// An array A consisting of N different integers is given. The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.

// Your goal is to find that missing element.

// Write a function:

// class Solution { public int solution(int[] A); }

// that, given an array A, returns the value of the missing element.

// For example, given array A such that:

//   A[0] = 2
//   A[1] = 3
//   A[2] = 1
//   A[3] = 5
// the function should return 4, as it is the missing element.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [0..100,000];
// the elements of A are all distinct;
// each element of array A is an integer within the range [1..(N + 1)]

func main() {
	// A := []int{2, 3, 1, 5}
	A := []int{2}

	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {
	if len(A) == 0 {
		return 1
	}

	out := make(map[int]bool, len(A)+1)

	for i := 1; i <= len(A)+1; i++ {
		out[i] = true
	}

	for _, v := range A {
		delete(out, v)
	}

	for r := range out {
		return r
	}

	return -1
}
