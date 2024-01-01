package main

import (
	"fmt"
	"math"
)

func main()  {
	var a float64
	var b float64
	fmt.Scan(&a, &b)
	c := math.Sqrt(a*a + b*b)
	fmt.Printf("%d", int(c))
}