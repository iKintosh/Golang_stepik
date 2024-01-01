package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	for idx, r := range str {
		if idx != 0 && idx != len(str) {
			fmt.Printf("*")
		}
		fmt.Printf("%c", r)
	}
}
