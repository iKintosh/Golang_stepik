package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	fmt.Scan(&a)
	fmt.Scan(&b)
	i := a + 1
	for ; i <= b; i++ {
		a = a + i
	}
	fmt.Println(a)
}