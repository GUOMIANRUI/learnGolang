package main

import "fmt"

func main() {
	// 布尔类型， 表示真假，标识符bool, 字面量true/false
	var zero bool

	isBoy := true
	isGirl := false

	fmt.Println(zero, isBoy, isGirl)
	// 操作
	// 逻辑运算(与 &&, 或 ||, 非 !)
	// aBool && bBool 全真才真
	fmt.Println("&&")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	// || 一真即真
	fmt.Println("||")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)
	// !aBool 取反
	fmt.Println("!")
	fmt.Println(!true)
	fmt.Println(!false)
	// 关系运算（==， !=）
	fmt.Println(isBoy == isGirl)
	fmt.Println(isBoy != zero)

	fmt.Printf("%T\n", isBoy)
}
