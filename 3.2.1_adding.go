package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func cleaner(x string) int64 {
	var collected_digits []rune
	for _, r := range x {
		if unicode.IsDigit(r) {
			collected_digits = append(collected_digits, r)
		}
	}
	num, err := strconv.ParseInt(string(collected_digits), 10, 64)
	if err != nil {
		return 0
	} else {
		return num
	}
}

func adding(a string, b string) int64  {
	first_num := cleaner(a)
	second_num := cleaner(b)
	return first_num + second_num
}

func main()  {
	var a = "%^80"
	var b = "hhhhh20&&&&nd"
	fmt.Print(adding(a, b))
}