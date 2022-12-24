package main

import "fmt"

func main(){
	var num_str string
	
	fmt.Scan(&num_str)
	if len(num_str) < 2{
		fmt.Println(0)
	} else {
		last_figure := num_str[len(num_str) - 2]
		fmt.Println(string(last_figure))
	}
}