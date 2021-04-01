package main

import "fmt"

func main() {

	// 闭包 有三要素，一是函数嵌套，二是内部函数引用了外部函数的变量（Base），三是外部函数又将内部函数作为返回值返回。 闭包使得局部变量在函数外被访问成为可能。
	// 如下 外函数的Base用完后就该销毁的 但是闭包使得其保留

	addBase := func(base int) func(int) int {
		// 返回函数
		return func(n int) int {
			return base + n
		}
	}
	add8 := addBase(8)
	fmt.Printf("%T\n", add8) // func(int) int
	fmt.Println(add8(10))

	fmt.Println(addBase(2)(3))
	fmt.Println(addBase(10)(3))
	fmt.Println(addBase(100)(3))
}
