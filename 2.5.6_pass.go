package main

import (
	"fmt"
	"unicode/utf8"
	"unicode"
)

func main() {
	var password string
	fmt.Scan(&password)
	if utf8.RuneCountInString(password) == len(password) && len(password) >= 5 {
		for _, r := range password {
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
				fmt.Printf("Wrong password")
				return
			}
		}
		fmt.Printf("Ok")
		return
	}
	fmt.Printf("Wrong password")
}
