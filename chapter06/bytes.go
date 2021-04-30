package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.NewBufferString("abcdefg")
	buffer.Write([]byte("123"))
	buffer.WriteString("!@#")

	fmt.Println(buffer.String())

	bytes := make([]byte, 2)
	buffer.Read(bytes)         // 一次读两个
	fmt.Println(string(bytes)) // ab

	line, err := buffer.ReadString('!') // 注意特殊符号用单引号   buffer.ReadBytes也可以
	fmt.Println(line, err)              // cdefg123! <nil>
}
