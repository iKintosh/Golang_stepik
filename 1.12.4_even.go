package main

import (
	"fmt"
)

func main() {
	var num_of_elements int
	fmt.Scan(&num_of_elements)
	slice_of_nums := make([]int, 0, 10)

	for i := 0; i < num_of_elements; i++ {
		var temp int
		fmt.Scan(&temp)
		slice_of_nums = append(slice_of_nums, temp)
	}

	for i := 0; i < len(slice_of_nums); i += 2 {
		fmt.Print(slice_of_nums[i], " ")
	}
}