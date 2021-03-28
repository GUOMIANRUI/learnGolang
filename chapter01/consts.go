package main

import "fmt"

func main() {
	const NAME string = "ll"
	// 常量不可更改
	fmt.Println(NAME)

	// 省略类型
	const USER = "guomianrui"
	//  定义多个常量（类型相同）
	const A, B, C int = 1, 2, 3
	// 定义多个常量（类型不相同）
	const (
		C3 string = "silence"
		C4 int    = 1
	)
	// 定义多个常量 省略类型
	const C5, C6 = "math", 3

	// 类型 值 都一样可以省略
	const (
		C7 int = 1
		C8
		C9
		C10 float64 = 3.14
		C11
		C12
		C13 string = "kk"
	)
	fmt.Println(C7, C8, C9)

	// 应用： 枚举， const + iota
	const (
		E1 int = iota
		E2
		E3
	)
	const (
		E4 int = iota
		E5
		E6
	)
	fmt.Println(E1, E2, E3)
	fmt.Println(E4, E5, E6)
	const (
		E7 int = (iota + 1) * 100
		E8
		E9
	)
	fmt.Println(E7, E8, E9)
}
