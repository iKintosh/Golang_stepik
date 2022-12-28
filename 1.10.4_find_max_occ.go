package main

import "fmt"

func main() {
	var (
		num     int
		max_num int = 0
		cnt     int = 0
	)

	for fmt.Scan(&num); num != 0; fmt.Scan(&num){
		if num > max_num {
			max_num = num
			cnt = 1
		} else if num == max_num {
			cnt = cnt + 1
		}
	}
	fmt.Println(cnt)
}
