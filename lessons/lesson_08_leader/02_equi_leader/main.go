package main

import "fmt"

// A non-empty array A consisting of N integers is given.

// The leader of this array is the value that occurs in more than half of the elements of A.

// An equi leader is an index S such that 0 ≤ S < N − 1 and two sequences A[0], A[1], ..., A[S] and A[S + 1], A[S + 2], ..., A[N − 1] have leaders of the same value.

// For example, given array A such that:

//     A[0] = 4
//     A[1] = 3
//     A[2] = 4
//     A[3] = 4
//     A[4] = 4
//     A[5] = 2
// we can find two equi leaders:

// 0, because sequences: (4) and (3, 4, 4, 4, 2) have the same leader, whose value is 4.
// 2, because sequences: (4, 3, 4) and (4, 4, 2) have the same leader, whose value is 4.
// The goal is to count the number of equi leaders.

// Write a function:

// func Solution(A []int) int

// that, given a non-empty array A consisting of N integers, returns the number of equi leaders.

// For example, given:

//     A[0] = 4
//     A[1] = 3
//     A[2] = 4
//     A[3] = 4
//     A[4] = 4
//     A[5] = 2
// the function should return 2, as explained above.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [1..100,000];
// each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].

func main() {
	A := []int{4, 3, 4, 4, 4, 2}

	result := Solution(A)
	fmt.Println("result:", result)
}

func Solution(A []int) int {

	// fmt.Println(A)

	// if len(A) == 0 {
	// 	return 0
	// }

	// if len(A) == 1 {
	// 	return 0
	// }

	// m := map[int]int{A[0]: 1} // map[value]count

	// for i := 1; i < len(A); i++ {
	// 	v := A[i]
	// 	if count, ok := m[v]; ok {
	// 		m[v] = count + 1
	// 	} else {
	// 		m[v] = 1
	// 	}

	// }

	// fmt.Println(m)
	// leader := A[0]
	// leaderCnt := 1
	// for k, v := range m {
	// 	if v > leaderCnt {
	// 		leader = k
	// 		leaderCnt = v
	// 	}
	// }

	// fmt.Println("leader:", leader, "leaderCnt:", leaderCnt)

	// for i := 0; i < len(A); i++ {

	// }

	return 0
}
