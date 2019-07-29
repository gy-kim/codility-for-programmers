package main

import "fmt"

// A non-empty array A consisting of N integers is given. A pair of integers (P, Q), such that 0 ≤ P < Q < N, is called a slice of array A (notice that the slice contains at least two elements). The average of a slice (P, Q) is the sum of A[P] + A[P + 1] + ... + A[Q] divided by the length of the slice. To be precise, the average equals (A[P] + A[P + 1] + ... + A[Q]) / (Q − P + 1).

// For example, array A such that:

//     A[0] = 4
//     A[1] = 2
//     A[2] = 2
//     A[3] = 5
//     A[4] = 1
//     A[5] = 5
//     A[6] = 8
// contains the following example slices:

// slice (1, 2), whose average is (2 + 2) / 2 = 2;
// slice (3, 4), whose average is (5 + 1) / 2 = 3;
// slice (1, 4), whose average is (2 + 2 + 5 + 1) / 4 = 2.5.
// The goal is to find the starting position of a slice whose average is minimal.

// Write a function:

// func Solution(A []int) int

// that, given a non-empty array A consisting of N integers, returns the starting position of the slice with the minimal average. If there is more than one slice with a minimal average, you should return the smallest starting position of such a slice.

// For example, given array A such that:

//     A[0] = 4
//     A[1] = 2
//     A[2] = 2
//     A[3] = 5
//     A[4] = 1
//     A[5] = 5
//     A[6] = 8
// the function should return 1, as explained above.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [2..100,000];
// each element of array A is an integer within the range [−10,000..10,000].

func main() {
	A := []int{4, 2, 2, 5, 1, 5, 8}

	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {

	if len(A) < 2 {
		return -1
	}
	if len(A) == 2 {
		return 0
	}

	minIdx := 0
	minAvg := float64(0)

	for i := 0; i < len(A)-1; i++ {
		avg2 := float64(A[i]+A[i+1]) / 2
		fmt.Println("i:", i, "avg2:", avg2)

		if i == 0 {
			minAvg = avg2
		}

		if avg2 < minAvg {
			minAvg = avg2
			minIdx = i
		}

		if i < len(A)-2 {
			avg3 := float64(A[i]+A[i+1]+A[i+2]) / 3
			fmt.Println("avg3:", avg3)

			if avg3 < minAvg {
				minAvg = avg3
				minIdx = i
			}
		}

	}

	return minIdx
}
