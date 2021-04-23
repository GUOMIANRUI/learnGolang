package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// https://golang.google.cn/pkg/os/#File
	path := "user.txt"
	// os.Open(path) // 会返回两个变量： 文件指针、错误信息
	file, err := os.Open(path)

	// fmt.Println(err)
	// fmt.Printf("%#v\n", file)
	if err != nil {
		fmt.Println(err)
	} else {
		// var bytes []byte = make([]byte, 20) // 定义字节切片
		bytes := make([]byte, 20)

		// n, err := file.Read(bytes)
		// fmt.Println(n, err)
		// // fmt.Println(bytes[:n]) 转换成字符串
		// fmt.Println(string(bytes[:n]))

		// n, err = file.Read(bytes)
		// fmt.Println(string(bytes[:n])) // 会接着上面继续读n个  读到末尾时返回EOF
		for {
			n, err := file.Read(bytes)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} else {
				fmt.Println(string(bytes[:n])) // Println 打印换行 Print 不打印换行
			}
		}
		file.Close()
	}

}
