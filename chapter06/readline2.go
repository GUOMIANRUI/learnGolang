package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("user.txt")
	if err == nil {
		defer file.Close()

		reader := bufio.NewReader(file) // 带缓冲的读  在内存里面开了一块空间做缓冲

		// bytes := make([]byte, 5)
		for {
			// line, err := reader.ReadSlice('\n')
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)

				}
				break
			} else {
				// fmt.Println(err, string(line))
				fmt.Println(err, line)
			}
		}
	}
}
