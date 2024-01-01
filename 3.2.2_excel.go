package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertNumber(input string) (float64, error) {
	converted_str := strings.Replace(input, " ", "", -1)
	converted_str = strings.Replace(converted_str, ",", ".", -1)
	if num, err := strconv.ParseFloat(converted_str, 64); err == nil {
		return num, nil
	} else {
		fmt.Println(err)
		return 0, err
	}
	strconv.FormatUint()
	even := []rune{}
}

func main()  {
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	numbers := strings.Split(string(line), ";")
	first_num, _ := convertNumber(numbers[0])
	second_num, _ := convertNumber(numbers[1])
	fmt.Printf("%.4f", first_num/second_num)
}