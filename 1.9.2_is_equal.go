package main

import "fmt"

func main(){
	var num string
	fmt.Scan(&num)
	switch {
	case num[0] == num[1]:
		fmt.Println("NO")
	case num[0] == num[2]:
		fmt.Println("NO")
	case num[1] == num[2]:
		fmt.Println("NO")
	default:
		fmt.Println("YES")
	}
}