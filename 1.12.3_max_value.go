package main
import "fmt"
import "math"

func main()  {
	array := [5]int{}
	var a int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}

	max_value := int(math.Inf(-1))
    for i := 0; i < len(array); i++ {
		if array[i] > max_value {
			max_value = array[i]
		}
	}
	fmt.Print(max_value)
}