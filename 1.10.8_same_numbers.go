package main

import "fmt"

func main() {
	var (
		first  string
		second string
	)
	fmt.Scan(&first)
	fmt.Scan(&second)

	for i := 0; i < len(first); i++ {
		letter_f := first[i]
		for j := 0; j < len(second); j++ {
			if letter_f == second[j] {
				fmt.Print(string(letter_f), " ")
			}
		}
	}
}
