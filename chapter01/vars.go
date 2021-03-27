package main

import "fmt"

var version string = "1.0"

func main() {
	// 定义一个string类型的变量me
	/*
		变量名需满足标识符命名规则：
		1. 必须由非空的unicode字符串组成、数字、_
		2. 不能以数字开头
		3. 不能为go的关键字（25个）

		4. 避免和go预定义标识符冲突，true/false
		5. 驼峰   myName
		6. 标识符区分大小写
	*/
	var me string

	me = "guomianrui"
	fmt.Println(me)
	fmt.Println(version)

	var name, user string = "guomianrui", "marin"
	fmt.Println(name, user)

	var (
		age    int     = 22
		height float64 = 1.78
	)

	fmt.Println(age, height)

	// 类型省略 会通过值推出类型  推荐var()这种方式写多个变量 函数外常用
	var (
		s = "kk"
		a = "22"
	)
	fmt.Println(s, a)

	// 这是一个简短声明  只能在函数内部使用   函数内常用
	isBoy := false
	fmt.Println(isBoy)
	var (
		ss     = "kk"
		aa int = 34
	)
	ss, aa, isBoy = "ll", 22, true

	fmt.Println(s, ss, aa)
	s, ss = ss, s
	fmt.Println(s, ss, aa)
}
