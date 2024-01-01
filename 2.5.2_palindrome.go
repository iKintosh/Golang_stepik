package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Scan(&word)
	chars := []rune(word)
	l := len(chars)
	for idx, r := range chars {
		if chars[l-idx-1] != r {
			fmt.Printf("Нет")
			return
		}
	}
	fmt.Printf("Палиндром")
}
