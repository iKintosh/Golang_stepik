package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	for _, r := range str {
		if strings.Count(str, string(r)) == 1 {
			fmt.Printf("%c", r)
		}
	}
}
