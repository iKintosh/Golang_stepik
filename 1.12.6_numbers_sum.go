package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var input_number string
	fmt.Scan(&input_number)
	var output int
	for i := 0; i < len(input_number); i++ {
		num, err := strconv.Atoi(string(input_number[i]))

		if err != nil {
			return
		}
		output += num
	}
	fmt.Print(output)
}