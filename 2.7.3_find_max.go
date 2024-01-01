package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputNumber string
	fmt.Scan(&inputNumber)
	maxValue := -10
	for _, r := range inputNumber {
		num, err := strconv.Atoi(string(r))

		if err != nil {
			return
		} else {
			if num > maxValue {
				maxValue = num
			}
		}

	}
	fmt.Print(maxValue)
}
