package main

import "fmt"

func main() {
	A := 2
	B := A
	B = 3
	fmt.Println(A, B)

	// 指针直接指向对象的内存地址 修改内存地址中的数据
	var cc *int = &A
	C := &A
	fmt.Printf("%T %T %p\n", C, cc, cc)

	fmt.Println(*cc) // 打印值
	*cc = 4
	fmt.Println(A)
}
