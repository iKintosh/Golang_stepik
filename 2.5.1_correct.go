package main

import (
	"bufio"
	"os"
	"unicode/utf8"
	"unicode"
	"fmt"
)

func main()  {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	firstChar, _ := utf8.DecodeRuneInString(text)
	lastChar, _ := utf8.DecodeLastRuneInString(text)
	if unicode.IsUpper(firstChar) && lastChar == '.' {
		fmt.Printf("Right")
		return
	}
	fmt.Printf("Wrong")
}