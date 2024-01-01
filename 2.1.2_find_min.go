package main

import "fmt"

func fibonacci(n int) int {
    if n == 1 {
        return 1
    } else if n == 2 {
        return 1
    }
    fib1 := 1
    fib2 := 1
    for i := 2; i < n; i++ {
        fib1, fib2 = fib2, fib1+fib2
    }
    return fib2
}

func sumInt(a ...int) (int, int) {
	numArgs := len(a)
	sumArgs := 0
	for _, value := range a {
		sumArgs += value
	}
	return numArgs, sumArgs
}




func minimumFromFour() int {
	min_value := 10000000
    for i := 0; i < 4; i++ {
		var temp int
		fmt.Scan(&temp)
		if temp < min_value {
			min_value = temp
		}
	}
	return min_value
}

func main()  {
	minimumFromFour()
}