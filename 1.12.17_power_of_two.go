package main

import (
	"fmt"
	"math"
)

func power_or_two(num int, pow int) int {
	float_num := float64(num)
	float_pow := float64(pow)
	res := math.Pow(float_num, float_pow)
	return int(res)
}

func main()  {
	var upper_bound int
	fmt.Scan(&upper_bound)
	var current_num int = 1 
	for i := 1; current_num <= upper_bound; i++ {
		fmt.Print(current_num, " ")
		current_num = power_or_two(2, i)
	}
	
}