package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Create("user3.txt")
	if err == nil {
		defer file.Close()

		writer := bufio.NewWriter(file)

		writer.WriteString("guomianrui")
		writer.Write([]byte("123456"))
		writer.Flush()
	}
}
