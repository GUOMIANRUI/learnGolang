package main

import "fmt"

func main() {
	var age = 30

	// 擅用打印语句排错  检查逻辑
	age = age + 1
	fmt.Println("明年：", age)
	age = age + 1
	fmt.Println("后年：", age)

	fmt.Print("不会换行")
	fmt.Print("不会换行")
	fmt.Println("会换行")
	fmt.Println("会换行")
	fmt.Printf("%T, %s, %T, %d\n", "KK", "12", age, age)
}
