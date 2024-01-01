package main

import "fmt"

func main()  {
	var input_slice []byte
	var input_str string

	fmt.Scan(&input_str)
	for i := len(input_str) - 1 ; i >= 0; i-- {
		input_slice = append(input_slice, input_str[i])
	}
	reversed_str := string(input_slice)
	fmt.Printf("%s", reversed_str)
}