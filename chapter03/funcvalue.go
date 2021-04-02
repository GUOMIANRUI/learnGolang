package main

import "fmt"

func changeInt(a int) {
	a = 100
}

// 可通过指针改int类型数据
func changeIntByPoint(p *int) {
	*p = 100
}

func changeSlice(s []int) {
	s[0] = 100
}

func main() {
	num := 1
	changeInt(num)
	fmt.Println(num) // 1

	changeIntByPoint(&num)
	fmt.Println(num) // 100

	nums := []int{1, 2, 3}
	changeSlice(nums)
	fmt.Println(nums) // [100 2 3]

}
