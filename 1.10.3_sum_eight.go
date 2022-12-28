package main

import "fmt"

func main() {
	var (
		length int
		num    int
		sum    int = 0
	)
	fmt.Scan(&length)

	for i := 0; i < length; i++ {
		fmt.Scan(&num)
		if num%8 == 0 && num/100 == 0 && num/10 > 0 {
			sum = sum + num
		}
	}
	fmt.Println(sum)
}
