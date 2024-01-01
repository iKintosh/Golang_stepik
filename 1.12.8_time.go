package main

import "fmt"

func main() {
	var seconds int
	fmt.Scan(&seconds)
	hours := seconds/3600
	seconds = seconds - hours*3600
	minutes := seconds/60

	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}