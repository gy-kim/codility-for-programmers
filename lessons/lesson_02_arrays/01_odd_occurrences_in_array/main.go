package main

import "fmt"

func main() {
	arr := []int{9, 3, 9, 3, 9, 7, 9}

	result := Solution(arr)
	fmt.Println("Result:", result)
}

func Solution(A []int) int {
	list := make(map[int]bool)
	result := 0

	for _, v := range A {
		if _, ok := list[v]; ok {
			delete(list, v)
			continue
		}
		result = v
		list[v] = true
	}

	return result
}
