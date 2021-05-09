package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new GOroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {
	// 创建一个go程 去执行newTask() 流程
	go newTask()

	// 主进程同时继续执行

	// fmt.Println("main go routine exit")
	i := 0

	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
	// main goroutine: i = 1
	// new GOroutine : i = 1
	// main goroutine: i = 2
	// new GOroutine : i = 2
	// main goroutine: i = 3
	// new GOroutine : i = 3
}
