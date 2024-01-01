package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type GlobalIds struct {
	Global_id int
}

func main() {
	file, err := os.Open("data-20190514T0100.json")
	if err != nil {
		fmt.Print(err)
	}
	dec := json.NewDecoder(file)
	var globalIds []GlobalIds
	dec.Decode(&globalIds)
	sum := 0
	for _, gi := range globalIds {
		sum = sum + gi.Global_id
	}
	fmt.Print(sum)
}
