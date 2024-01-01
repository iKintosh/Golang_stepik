package main

import "fmt"

func main() {
	var lower_bound int
	var higher_bound int
	var candidate int
	fmt.Scan(&lower_bound, &higher_bound)
	if higher_bound < 0 {
		candidate = (int(higher_bound/7) - 1) * 7
	} else {
		candidate = int(higher_bound/7) * 7
	}

	if candidate >= lower_bound {
		fmt.Printf("%d", candidate)
		return
	}
	fmt.Printf("NO")
}
