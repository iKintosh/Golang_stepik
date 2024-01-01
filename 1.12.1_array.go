package main

import "fmt"

func main() {
	var workArray [10]int
	for i := 0; i < 10; i++ {
		_, err := fmt.Scan(&workArray[i])
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}
	for i := 0; i < 3; i++ {
		var num_1 int
		var num_2 int
		fmt.Scan(&num_1)
		fmt.Scan(&num_2)
		temp := workArray[num_1]
		workArray[num_1] = workArray[num_2]
		workArray[num_2] = temp
	}
	for i := 0; i<len(workArray); i++ {
		fmt.Printf("%d ", workArray[i])
	}
}