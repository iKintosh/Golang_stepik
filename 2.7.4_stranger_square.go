package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputNumber string
	fmt.Scan(&inputNumber)
	for _, r := range inputNumber {
		num, err := strconv.Atoi(string(r))

		if err != nil {
			return
		} else {
			fmt.Printf("%d", num*num)
		}

	}
}
