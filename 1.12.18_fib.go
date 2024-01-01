package main

import (
	"fmt"
)

func main()  {
	var seek int
	fmt.Scan(&seek)
	fib1, fib2 := 1, 1
	cnt := 2
	for ; seek > fib2; {
		fib1, fib2 = fib2, fib1+fib2
		cnt++
	}
	if seek == fib2 {
		fmt.Print(cnt)
		return
	}
	fmt.Print(-1)
}