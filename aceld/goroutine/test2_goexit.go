package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		//用go创建承载一个形参为空，返回值为空的一个函数
		go func() {
			defer fmt.Println("A.defer")

			func() {
				defer fmt.Println("B.defer")
				runtime.GOexit() //退出当前goroutine
				fmt.Println("B")
			}() // 后面加了() 表示调用 防止函数创建了没调用

			fmt.Println("A")
		}()
	*/

	go func(a int, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true // 补： 在这里不能通过flag := go func ...return 获得子go程的返回值  需要通过后面学的管道  channel
	}(10, 20)

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
