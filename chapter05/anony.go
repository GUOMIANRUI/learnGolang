package main

import "fmt"

func main() { // 匿名结构体，不定义结构体的名字 直接用结构体声明一个变量
	var me struct { // 定义变量 struct定义结构体
		ID   int
		Name string
	}

	fmt.Printf("%T\n", me) // 一个没有名字的结构体
	fmt.Printf("%#v\n", me)
	fmt.Println(me.ID)
	me.Name = "one"
	fmt.Printf("%#v\n", me)

	// 用字面量 简短声明 初始化匿名结构体
	me2 := struct {
		ID   int
		Name string
	}{1, "guo"}

	fmt.Printf("%#v\n", me2)

}
