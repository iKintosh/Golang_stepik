package main
import "fmt"
import "math"

func main()  {
	var N int
	var slice []int
	var a int
	min_value := math.MaxInt
	cnt := 0

	fmt.Scan(&N)
	for i:=0; i < N; i++{
		fmt.Scan(&a)
		slice = append(slice, a)
		if a < min_value {
			min_value = a
		}

	}

	
    for i := 0; i < len(slice); i++ {
		if slice[i] == min_value {
			cnt++
		}
	}
	fmt.Print(cnt)
}