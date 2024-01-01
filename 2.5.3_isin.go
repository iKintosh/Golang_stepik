package main

import (
	"fmt"
	"strings"
)

func main() {
	var parent string
	var child string
	fmt.Scan(&parent)
	fmt.Scan(&child)
	fmt.Printf("%d", strings.Index(parent, child))
}
