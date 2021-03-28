package main

import "fmt"

func main() {
	// 整数类型
	// 标识符： int/int*/uintptr/byte/runc
	// 字面量: 十进制, 八进制 0777 = 7 * 8^2 + 7 * 8^1 + 7 * 8^0 ，十六进制0X0-9A-F
	var age int = 31
	fmt.Printf("%T %d\n", age, age)

	fmt.Println(0777, 0X10)

	// 操作
	// 算数运算（+ - * / %, ++, --）
	fmt.Println(1 + 2)
	fmt.Println(3 - 10)
	fmt.Println(3 * 9)
	fmt.Println(9 / 2)
	fmt.Println(9 % 2)

	age++
	fmt.Println(age)
	age--
	fmt.Println(age)

	// 赋值运算

}
