package main

import "fmt"

func main() {
	var (
		a int
		b int
		c int
	)
	fmt.Scan(&a, &b, &c)
	if c*c == a*a + b*b {
		fmt.Printf("Прямоугольный")
		return
	}
	fmt.Printf("Непрямоугольный")
	
}