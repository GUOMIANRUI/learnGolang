package main

import "fmt"

func main() { // 匿名结构体，不定义结构体的名字
	var me struct { // 定义变量 struct定义结构体
		ID   int
		Name string
	}

	fmt.Printf("%T\n", me) // 一个没有名字的结构体
	fmt.Printf("%#v\n", me)
}
