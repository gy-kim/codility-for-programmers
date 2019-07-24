package main

import "fmt"

func main() {
	A := []int{3, 8, 9, 7, 6}
	// {6,3,8,9,7}
	K := 6

	result := Solution(A, K)
	fmt.Println("result:", result)
}
func Solution(A []int, K int) []int {
	if len(A) == 0 || K == 0 || len(A) == K {
		return A
	}
	if len(A) > K {
		div := len(A) % K
		return append(A[div:], A[:div]...)
	} else {
		div := K % len(A)
		return append(A[len(A)-div:], A[:len(A)-div]...)
	}
}
