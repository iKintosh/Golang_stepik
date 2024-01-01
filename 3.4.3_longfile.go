package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("task.data")
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	cnt := 1
	for {
		v, err := r.ReadSlice(';')
		if err != nil {
			fmt.Print(err)
			return
		}
		if string(v) == "0;" {
			fmt.Print(cnt)
			return
		}
		cnt++
	}
	fmt.Print(cnt)
}
