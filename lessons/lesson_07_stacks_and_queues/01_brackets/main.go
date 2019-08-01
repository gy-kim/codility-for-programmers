package main

import "fmt"

// A string S consisting of N characters is considered to be properly nested if any of the following conditions is true:

// S is empty;
// S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
// S has the form "VW" where V and W are properly nested strings.
// For example, the string "{[()()]}" is properly nested but "([)()]" is not.

// Write a function:

// class Solution { public int solution(String S); }

// that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.

// For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [0..200,000];
// string S consists only of the following characters: "(", "{", "[", "]", "}" and/or ")".

func main() {
	S := "{[()()]}"

	result := Solution(S)
	fmt.Println("result:", result)
}

func Solution(S string) int {

	if len(S) == 0 {
		return 1
	}

	chars := map[string]int{"(": -3, "{": -2, "[": -1, "]": 1, "}": 2, ")": 3}

	stack := []int{}

	for _, c := range S {

		fmt.Println("c:", string(c), "stack:", stack)
		ch := string(c)
		num := chars[ch]

		if num < 0 {
			stack = append(stack, num)
		} else {
			if len(stack) == 0 {
				return 0
			}

			last := stack[len(stack)-1]

			if (last + num) == 0 {
				stack = stack[:len(stack)-1]
			} else {
				return 0
			}
		}

	}

	if len(stack) != 0 {
		return 0
	}

	return 1
}

// https://app.codility.com/demo/results/training4FMEPA-ZMA/
