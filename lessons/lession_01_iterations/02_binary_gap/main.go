package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := Solution(32)
	fmt.Println("result:", result)
}

func Solution(N int) int {
	i := int64(N)
	str := fmt.Sprint(strconv.FormatInt(i, 2))
	max, len := 0, 0
	for _, s := range str {
		if string(s) == "0" {
			len++
			continue
		}
		if max < len {
			max = len
			len = 0
		}
	}

	return max
}
