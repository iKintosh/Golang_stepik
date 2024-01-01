package main

import "fmt" // пакет используется для проверки ответа, не удаляйте его
import "strings"


type Battery struct {
	capacity string
}

func (b Battery) String() string {
	// var cpt int
    var capacity_str string = ""
    capacity_str = capacity_str + "["
	cpt := strings.Count(b.capacity, "1")
    uncharged := strings.Repeat(" ", 10-cpt)
    charged := strings.Repeat("X", cpt)
    capacity_str = capacity_str + uncharged
    capacity_str = capacity_str + charged
    capacity_str = capacity_str + "]"
    return capacity_str
}

func main()  {
	var bat string
	fmt.Scan(&bat)
	batteryForTest := Battery{capacity: bat}
	fmt.Print(batteryForTest)
}