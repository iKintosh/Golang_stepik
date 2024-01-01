package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Answer struct {
	Average float64
}

func main() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	input := ""
	for r.Scan() {
		tmp := r.Text()
		input = input + tmp
	}
	var json_data Group
	err := json.Unmarshal([]byte(input), &json_data)
	if err != nil {
		fmt.Print(err)
	}
	cntStudents := 0
	cntRating := 0
	for _, s := range json_data.Students {
		cntStudents++
		cntRating = cntRating + len(s.Rating)
	}
	ans := Answer{Average: float64(cntRating) / float64(cntStudents)}
	data, err := json.MarshalIndent(ans, "", "    ")
	if err != nil {
		fmt.Print(err)
	}
	w.WriteString(string(data))
	w.Flush()
}
