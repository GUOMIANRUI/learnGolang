package main

import (
	"fmt"

	"github.com/GUOMIANRUI/testmod/gopkg" // 调用本地时  要把gomod的前缀加上
	"github.com/howeyc/gopass"
)

func main() {
	fmt.Println("请输入密码：")
	if bytes, err := gopass.GetPasswd(); err == nil {
		fmt.Println(string(bytes))
	}
	fmt.Println(gopkg.Version)
}
