package main

import "fmt"

func addN(n int) int {
	if n == 1 {
		return 1
	}
	// fmt.Println(n)
	return n + addN(n-1)
}

// n阶乘:
// n! = 1 * 2 * 3 * 4 * 5       n = 0 -->  n != 1
func Factorial(n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return -1 // 小于0的时候 返回-1表示错误
	}
	return n * Factorial(n-1)
}

func main() {
	fmt.Println(addN(5))
	fmt.Println(Factorial(10))
}
