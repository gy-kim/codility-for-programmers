package main

import "fmt"

func fnFactorial(n int) {
	factorial := 1
	for i := 1; i <= n; i++ {
		fmt.Print(i, " ")
		factorial *= i
	}
	fmt.Println("\nfactorial:", factorial)
}

func fnAsterisks(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func fnAsterisks2(n int) {
	for i := n; i > 0; i-- {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func fnFibonacci(n int) {
	a := 0
	b := 1
	for {
		if a > n {
			break
		}
		fmt.Print(a, " ")
		c := a + b
		a = b
		b = c
	}
}

func main() {
	// fnFactorial(5)
	// fnAsterisks(5)
	// fnAsterisks2(4)
	fnFibonacci(15)

}
