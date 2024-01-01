package main

import "fmt"

func main() {
	var (
		a int
		b int
		c int
	)
	fmt.Scan(&a, &b, &c)
	if a + b > c && a + c > b && b + c > a {
		fmt.Printf("Существует")
		return
	}
	fmt.Printf("Не существует")
	
}

