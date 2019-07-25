package main

import "fmt"

// A non-empty array A consisting of N integers is given. Array A represents numbers on a tape.

// Any integer P, such that 0 < P < N, splits this tape into two non-empty parts: A[0], A[1], ..., A[P − 1] and A[P], A[P + 1], ..., A[N − 1].

// The difference between the two parts is the value of: |(A[0] + A[1] + ... + A[P − 1]) − (A[P] + A[P + 1] + ... + A[N − 1])|

// In other words, it is the absolute difference between the sum of the first part and the sum of the second part.

// For example, consider array A such that:

//   A[0] = 3
//   A[1] = 1
//   A[2] = 2
//   A[3] = 4
//   A[4] = 3
// We can split this tape in four places:

// P = 1, difference = |3 − 10| = 7
// P = 2, difference = |4 − 9| = 5
// P = 3, difference = |6 − 7| = 1
// P = 4, difference = |10 − 3| = 7
// Write a function:

// func Solution(A []int) int

// that, given a non-empty array A of N integers, returns the minimal difference that can be achieved.

// For example, given:

//   A[0] = 3
//   A[1] = 1
//   A[2] = 2
//   A[3] = 4
//   A[4] = 3
// the function should return 1, as explained above.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [2..100,000];
// each element of array A is an integer within the range [−1,000..1,000].

func main() {
	A := []int{3, 1, 2, 4, 3}

	result := Solution(A)
	fmt.Println("Solution:", result)
}

func Solution(A []int) int {

	min := 0
	list := make(map[int]int)

	for i := 0; i < len(A); i++ {
		if i == 0 {
			list[i] = A[i]
			continue
		}
		list[i] = list[i-1] + A[i]
	}
	// fmt.Println(list)

	for i := 0; i < len(list)-1; i++ {
		left := list[i]
		right := list[len(list)-1] - left
		result := left - right
		if result < 0 {
			result *= -1
		}
		if i == 0 {
			min = result
			continue
		}
		if min > result {
			min = result
		}
		// fmt.Println("left:", left, "right:", right, "result:", result, "min:", min)
	}

	return min
}
