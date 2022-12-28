package main

import "fmt"

func main() {
	var (
		account int
		per int
		goal int
		intermediate_accont int
	)
	fmt.Scan(&account)
	fmt.Scan(&per)
	fmt.Scan(&goal)

	intermediate_accont = account



	for i := 1; ; i++ {
		intermediate_accont = intermediate_accont * (100 + per) / 100
		if intermediate_accont >= goal {
			fmt.Println(i)
			break
		}
	}
}
