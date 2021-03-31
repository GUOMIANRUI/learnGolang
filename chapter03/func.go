package main

import "fmt"

// 定义 (无参， 无返回值)
func sayHelloWorld() {
	fmt.Println("Hello World!!!")
}

// 定义有参数(形参)
func sayHi(name string) {
	fmt.Println("你好：", name)
}

// 有参 有返回值
func add(a int, b int) int {
	return a + b
}

func main() {
	// 调用函数
	sayHelloWorld()

	// 打印标识符
	fmt.Printf("%T\n", sayHelloWorld)

	sayHi("棉锐") // 实参

	rt := add(6, 6)
	fmt.Println(rt)
}
