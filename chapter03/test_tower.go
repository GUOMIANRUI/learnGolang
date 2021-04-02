package main

import "fmt"

// 1层：
// A -> C
// 2层：
// A -> B
// A -> C
// B -> C

// a 源  b 借助的 c目的
func tower(a, b, c string, layer int) {
	if layer == 1 {
		fmt.Println(a, "->", c)
		return
	}
	// a n-1 借助c 移动到b上
	tower(a, c, b, layer-1)
	fmt.Println(a, "->", c)
	// b n-1 借助a 移动到c上
	tower(b, a, c, layer-1)
}

func main() {
	fmt.Println("1层：")
	tower("A", "B", "C", 1)
	fmt.Println("2层：")
	tower("A", "B", "C", 2)
	fmt.Println("3层：")
	tower("A", "B", "C", 3)
	fmt.Println("4层：")
	tower("A", "B", "C", 4)
}
