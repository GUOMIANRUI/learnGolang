package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("user.txt") // 不用自己关闭
	if err == nil {
		fmt.Println(string(bytes))
	}
}
