package main

import "fmt"

func test(x1 *int, x2 *int) {
    *x1, *x2 = *x2, *x1
    fmt.Printf("%d %d", *x1, *x2)
}

func main()  {
	x1 := 1
	x2 := 2
	test(&x1, &x2)
	fmt.Printf("%d %d", x1, x2)
}
