package main

import "fmt"

// A string S consisting of N characters is called properly nested if:

// S is empty;
// S has the form "(U)" where U is a properly nested string;
// S has the form "VW" where V and W are properly nested strings.
// For example, string "(()(())())" is properly nested but string "())" isn't.

// Write a function:

// func Solution(S string) int

// that, given a string S consisting of N characters, returns 1 if string S is properly nested and 0 otherwise.

// For example, given S = "(()(())())", the function should return 1 and given S = "())", the function should return 0, as explained above.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [0..1,000,000];
// string S consists only of the characters "(" and/or ")".

func main() {
	S := "(()(())())"

	result := Solution(S)

	fmt.Println("result:", result)
}

func Solution(S string) int {

	if len(S) == 0 {
		return 1
	}

	arr := map[string]int{"(": -1, ")": 1, "/": 0}

	qu := []int{}

	for _, c := range S {
		ch := string(c)
		if num := arr[ch]; num == 0 {
			continue
		} else if num < 0 {
			qu = append(qu, num)
			continue
		} else {
			if len(qu) == 0 {
				return 0
			}
			last := qu[len(qu)-1]
			if last+num != 0 {
				return 0
			}
			qu = qu[:len(qu)-1]
		}

		fmt.Println(qu)
	}

	if len(qu) != 0 {
		return 0
	}
	return 1
}

// https://app.codility.com/demo/results/trainingHB2BA9-V87/
