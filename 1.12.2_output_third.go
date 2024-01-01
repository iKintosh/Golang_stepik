package main
import "fmt"

func main() {
    var number_of_elements uint
	fmt.Scan(&number_of_elements)
	slice_of_nums := make([]int,0,10) 

	for i := 0; i < int(number_of_elements); i++ {
		var temp int
		fmt.Scan(&temp)
		slice_of_nums = append(slice_of_nums, temp)
	}
	fmt.Print(slice_of_nums[3])
}