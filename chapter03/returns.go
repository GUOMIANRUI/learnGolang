package main

import "fmt"

// func testReturn() int {
// 	fmt.Println("return 1")
// 	return 1 // 碰到return 就返回
// }

func calc(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

// 返回值带参数 默认值是0 代码量多、逻辑复杂的时候不介意用   应该用上面那种显式的
func calc2(a, b int) (sum int, diff int, product int, quotient int) {
	sum = a + b
	diff = a - b
	product = a * b
	quotient = a / b
	return
}

func main() {
	a, b, c, d := calc(9, 3)
	// a, b, _, _ := calc(9, 3) // 不关心的可以用_
	fmt.Println(a, b, c, d)
	fmt.Println(calc2(10, 2))
}
