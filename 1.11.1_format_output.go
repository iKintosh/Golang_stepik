package main

import "fmt"
import "os"
import "math"

func main() {
	var input_num float64
	var output_value float64
	fmt.Scan(&input_num)
	if input_num <= 0 {
		fmt.Printf("число %2.2f не подходит", input_num)
		os.Exit(0)
	} else if input_num >= 10000 {
		fmt.Printf("%e", input_num)
		os.Exit(0)
	}
	output_value = math.Pow(input_num, 2)
	fmt.Printf("%.4f", output_value)
}