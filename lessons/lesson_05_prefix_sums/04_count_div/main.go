package main

import "fmt"

// Write a function:

// func Solution(A int, B int, K int) int

// that, given three integers A, B and K, returns the number of integers within the range [A..B] that are divisible by K, i.e.:

// { i : A ≤ i ≤ B, i mod K = 0 }

// For example, for A = 6, B = 11 and K = 2, your function should return 3, because there are three numbers divisible by 2 within the range [6..11], namely 6, 8 and 10.

// Write an efficient algorithm for the following assumptions:

// A and B are integers within the range [0..2,000,000,000];
// K is an integer within the range [1..2,000,000,000];
// A ≤ B.

func main() {
	A := 6
	B := 11
	K := 2

	// A := 11
	// B := 345
	// K := 17

	result := Solution(A, B, K)
	fmt.Println("result:", result)
}

func Solution(A int, B int, K int) int {
	if A == 0 {
		return B/K + 1
	}
	return B/K - (A-1)/K // (B/K+1) - ((A-1)/K+1)
}
