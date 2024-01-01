package main

import (
	"fmt"
)

func main() {
	var str string
	var filtered []rune
	fmt.Scan(&str)
	for idx, r := range str {
		if idx % 2 != 0 {
			filtered = append(filtered, r)
		}
	}

	fmt.Printf("%s", string(filtered))
}
