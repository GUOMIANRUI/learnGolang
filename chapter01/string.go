package main

import "fmt"

func main() {
	// "" => 可解释
	// `` => 原生字符串
	// 特殊字符 \r \n \f \t \b \v
	var name string = "guo\tmianrui"
	var desc string = `我来自汕头`

	fmt.Println(name)
	fmt.Println("guo\\mianrui")
	fmt.Println(desc)

	// 操作
	// 算数运算符： + （连接）
	fmt.Println("我叫" + "guomianrui")
	// 关系运算（== != > >= < <=）

	fmt.Println("ab" == "bb")
	fmt.Println("ab" != "bb")
	fmt.Println("ab" < "bb") // 从左往右比较ascll码
	fmt.Println("ab" <= "bb")
	fmt.Println("ab" > "bb")
	fmt.Println("ab" >= "bb")
	fmt.Println("bb" >= "b")

	// 赋值
	s := "我叫"
	s += "guomianrui"

	// 字符串定义内容必须只能为ascll  字母啥的
	// 索引 0 - n-1 (n 字符串长度)
	desc = "abcdef"
	fmt.Printf("%T %c\n", desc[0], desc[0])
	// 切片[start:end] start - end-1
	fmt.Printf("%T %s\n", desc[0:2], desc[0:2])

	fmt.Println(len(desc))
}
