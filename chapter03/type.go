package main

import "fmt"

func add(a, b int) int {
	return a + b
}

// 函数也可以作为参数传递给另一个函数
// 传一个格式化的函数 callback ：将传递的数据按照每行打印还是按照一行按|分割打印
func print(callback func(...string), args ...string) { // 定义了一个 函数类型 的形参 名称叫callback    函数也可以传给函数
	fmt.Println("print函数输出：")
	callback(args...)
}

func list(args ...string) {
	for i, v := range args {
		fmt.Println(i, ":", v)
	}
}

func main() {
	fmt.Printf("%T\n", add) // func(int, int) int

	f := add

	fmt.Printf("%T\n", f) // func(int, int) int
	fmt.Println(f(1, 4))
	// 说明函数也可以赋值给变量

	print(list, "A", "C", "E")

	// 匿名函数
	sayHello := func(name string) {
		fmt.Println("Hello ", name)
	}
	sayHello("guomianrui")
	// 只调用一次可以这样写
	func(name string) {
		fmt.Println("Hi", name)
	}("marin")

	// 直接在传的时候定义一个匿名函数
	print(func(args ...string) {
		for _, v := range args {
			fmt.Println(v, "\t")
		}
		fmt.Println()
	}, "A", "C", "E")

}
