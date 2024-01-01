package main

import (
	"fmt"
	"strconv"
	"os"
)

func sum_str_values(num string) string {
	result := 0
	for i := 0; i < len(num); i++ {
		num, err := strconv.Atoi(string(num[i]))

		if err != nil {
			os.Exit(1)
		}
		result += num
	}
	string_result := strconv.Itoa(result)
	return string_result
}

func main()  {
	var input string
	fmt.Scan(&input)
	result := input

	for ; len(result) > 1; {
		result = sum_str_values(result)
	}

	fmt.Printf("%s", result)
}