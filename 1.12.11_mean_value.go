package main

import "fmt"

func main() {
	var (
		a float32
		b float32
	)
	fmt.Scan(&a, &b)
	mean_value := (a+b)/2
	fmt.Print(mean_value)
	
}