package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	path := "user.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close() // 没有错误的时候延迟文件关闭
		bytes, err := ioutil.ReadAll(file)
		fmt.Println(bytes, err)

	}
}
