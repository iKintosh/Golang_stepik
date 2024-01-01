package main

import "fmt"

func main()  {
	var korov_num uint8
	fmt.Scan(&korov_num)

	switch {
	case korov_num == 1:
		fmt.Printf("%d korova", korov_num)
	case korov_num > 1 && korov_num < 5:
		fmt.Printf("%d korovy", korov_num)
	case korov_num >= 5 && korov_num <= 20:
		fmt.Printf("%d korov", korov_num)
	case korov_num % 10 == 1:
		fmt.Printf("%d korova", korov_num)
	case korov_num % 10 > 1 && korov_num % 10 < 5:
		fmt.Printf("%d korovy", korov_num)
	case korov_num % 10 >= 5 && korov_num % 10 <= 9:
		fmt.Printf("%d korov", korov_num)
	case korov_num % 10 == 0 :
		fmt.Printf("%d korov", korov_num)
	}
	
}