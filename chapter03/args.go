package main

import "fmt"

func add(a, b int) int {
	return a + b
}

// 可变参数
func addN(a, b int, args ...int) int {
	// fmt.Println(a, b, args)
	// fmt.Printf("%T\n", args) // 切片类型
	// return 0
	total := a + b
	for _, v := range args {
		total += v
	}
	return total
}

func calc(op string, a, b int, args ...int) int {
	switch op {
	case "add":
		return addN(a, b, args...) // 用 ... 对切片进行解包 解成一个个int
	}
	return -1
}

func main() {
	fmt.Println(add(1, 5))
	fmt.Println(addN(1, 4))
	fmt.Println(addN(1, 4, 6))
	fmt.Println(addN(1, 4, 5, 3))

	fmt.Println(calc("add", 1, 2))
	fmt.Println(calc("add", 1, 2, 5))
	fmt.Println(calc("add", 1, 2, 5, 7))

	args01 := []int{1, 2, 5, 6, 10}
	fmt.Println(addN(1, 2, args01...))
	fmt.Println(calc("add", 1, 3, args01...))

	// 数组删除中间元素思路  将后半部分解包append到前面
	nums := []int{1, 3, 5, 8}
	nums = append(nums[:1], nums[2:]...)
	fmt.Println(nums)
}
