package main

import (
	"fmt"
)

func main()  {
	var input string
	var num_to_remove byte
	fmt.Scan(&input)
	fmt.Scanf("%c\n", &num_to_remove)

	for i := 0; i < len(input); i++ {
		if input[i] != num_to_remove {
			fmt.Printf("%c", input[i])
		}
	}
}