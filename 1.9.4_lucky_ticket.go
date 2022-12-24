package main

import (
	"fmt"
	"strconv"
)

func dti(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func main(){
	var ticket_num string
	fmt.Scan(&ticket_num)

	first_half := dti(string(ticket_num[0])) + dti(string(ticket_num[1])) + dti(string(ticket_num[2]))
	second_half := dti(string(ticket_num[3])) + dti(string(ticket_num[4])) + dti(string(ticket_num[5])) 

	if first_half == second_half {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}