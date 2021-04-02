package main

import "fmt"

func main() {
	// defer 延迟函数 在函数退出的时候执行 可以有多个defer 按照栈的先进后出顺序执行
	defer func() {
		fmt.Println("defer01")
	}()
	defer func() {
		fmt.Println("defer02")
	}()

	fmt.Println("main over")
}
