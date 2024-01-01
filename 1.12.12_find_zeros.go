package main

import "fmt"

func main() {
	var num_of_elements int
	fmt.Scan(&num_of_elements)
	var cnt int


	for i := 0; i < num_of_elements; i++ {
		var temp int
		fmt.Scan(&temp)
		if temp == 0 {
			cnt++
		}
	}
	fmt.Print(cnt)
}