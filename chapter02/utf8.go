package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "今天是个好天气"
	fmt.Println(len(s))                    // 21
	fmt.Println(utf8.RuneCountInString(s)) // 7
}
