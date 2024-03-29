package main

import "fmt"

// An array A consisting of N integers is given. The dominator of array A is the value that occurs in more than half of the elements of A.

// For example, consider array A such that

//  A[0] = 3    A[1] = 4    A[2] =  3
//  A[3] = 2    A[4] = 3    A[5] = -1
//  A[6] = 3    A[7] = 3
// The dominator of A is 3 because it occurs in 5 out of 8 elements of A (namely in those with indices 0, 2, 4, 6 and 7) and 5 is more than a half of 8.

// Write a function

// func Solution(A []int) int

// that, given an array A consisting of N integers, returns index of any element of array A in which the dominator of A occurs. The function should return −1 if array A does not have a dominator.

// For example, given array A such that

//  A[0] = 3    A[1] = 4    A[2] =  3
//  A[3] = 2    A[4] = 3    A[5] = -1
//  A[6] = 3    A[7] = 3
// the function may return 0, 2, 4, 6 or 7, as explained above.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [0..100,000];
// each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].

func main() {
	A := []int{3, 4, 3, 2, 3, -1, 3, 3}
	// A := []int{1, 2, 1}

	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {

	size := len(A)

	if size == 0 {
		return -1
	}

	if size == 1 {
		return 0
	}

	condition := (size / 2) + 1

	m := map[int]int{A[0]: 1} // value : count

	for i := 1; i < size; i++ {
		v := A[i]

		if cnt, ok := m[v]; ok {
			cnt++
			m[v] = cnt
		} else {
			m[v] = 1
		}

		if m[v] >= condition {
			return i
		}
	}

	return -1
}

// https://app.codility.com/demo/results/trainingFNVMHM-AFX/
