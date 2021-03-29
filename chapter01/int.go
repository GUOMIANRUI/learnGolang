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

	// 关系运算(== != >= <= > <)
	fmt.Println(2 == 3)
	fmt.Println(3 != 4)
	fmt.Println(2 > 3)
	fmt.Println(2 < 3)
	fmt.Println(2 <= 3)

	// 位运算 二进制的运算  10 => 2 (了解即可)
	// & | ^(异或) << >> &^(位清空)
	// 十进制   eg:7    7/2 1    3/2  1    1/2  1   0  0111
	// 2 => 0010
	// 7 & 2 => 0111 & 0010 => 0010  => 2
	// 7 | 2 => 0111 | 0010 => 0111  => 7
	// 7 ^ 2 => 0111 ^ 0010 => 0101  => 5
	// 2 << 1 => 0010 << 1  => 0100  => 4
	// 2 >> 1 => 0010 >> 1  => 0001  => 1
	// 7 &^ 2 => 0111 &^ 0010 => 0101  => 5

	fmt.Println(7 & 2)
	fmt.Println(7 | 2)

	// 赋值运算(=, +=, -=, *=, /=, %=, &=， |=，^=, <<=, >>=, &^= )
	age = 1
	age += 3 // age = age + 3
	fmt.Println(age)

	// int/uint/byte/rune/int*
	var intA int = 10
	var uintB uint = 3
	// 类型转换 转换成相同类型才可以进行计算
	fmt.Println(intA + int(uintB))
	fmt.Println(uint(intA) + uintB)

	// 位溢出   从大往小转可能出现溢出
	// var intC int = 0XFFFF
	// fmt.Println(intC, uint8(intC), int(intC))

	// fmt.Printf
	// int / uint / int* /uint*
	// byte, rune  这两个类型要用 单引号 引起来             实际上是uint8  int32
	var a byte = 'A'
	var w rune = '面'
	fmt.Println(a, w)

	fmt.Printf("%T %d %b %o %x\n", age, age, age, age, age) // %b 二进制 %o 八进制 %x 十六进制
	fmt.Printf("%T %c \n", a, a)
	fmt.Printf("%T %q %U \n", w, w, w)
	fmt.Printf("%5d\n", age)  // 占位符
	fmt.Printf("%05d\n", age) // 占位符 0
	fmt.Printf("%-5d\n", age) // 占位符 左对齐
}
