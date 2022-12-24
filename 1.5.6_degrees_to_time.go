package main

import "fmt"

func main() {
	var initial_minutes uint16
	fmt.Scan(&initial_minutes)
	hours := initial_minutes / 30
	minutes := (initial_minutes - hours*30) * 2
	fmt.Println("It is", hours, "hours", minutes, "minutes.")
}
