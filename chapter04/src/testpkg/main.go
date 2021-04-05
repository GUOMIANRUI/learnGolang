package main

import (
	"fmt"

	// lgopkg "gopkg" // 起别名防止包名冲突

	"github.com/GUOMIANRUI/gopkg"
	"github.com/howeyc/gopass"
)

func main() {
	// fmt.Println(lgopkg.Version)
	fmt.Println(gopkg.VERSION)
	// gopkg.PrintlnName()

	fmt.Println("请输入密码：")
	if bytes, err := gopass.GetPasswd(); err == nil {
		fmt.Println(string(bytes))
	}
}
