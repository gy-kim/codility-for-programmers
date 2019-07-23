package main

import (
	"fmt"
	"strconv"
)

func main() {

	i := int64(9)
	str := fmt.Sprint(strconv.FormatInt(i, 2))
	fmt.Println(str)
	for _, s := range str {
		fmt.Println(string(s))
	}
}
