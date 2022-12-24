package main

import "fmt"

func main(){
	var num_str string
	
	fmt.Scan(&num_str)
	last_figure := num_str[len(num_str) - 1]
	fmt.Println(string(last_figure))
}