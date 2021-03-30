package main

import "fmt"

func main() {
	var nums [10]int
	var t2 [5]bool
	var t3 [3]string

	fmt.Printf("%T %d", nums, nums)
	fmt.Println(t2)
	fmt.Printf("%q", t3)

	// 字面量
	nums = [10]int{10, 20, 30}
	fmt.Println(nums)

	nums = [10]int{0: 10, 9: 20}
	fmt.Println(nums)
	nums = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums)

	nums02 := [10]int{10, 20, 30}
	fmt.Printf("%T %#v\n", nums02, nums02)
	nums03 := [...]int{1, 2}
	fmt.Printf("%T %#v\n", nums03, nums03)
	nums04 := [15]int{1: 10, 5: 20, 14: 30}
	fmt.Printf("%T %#v\n", nums04, nums04)

	// 操作
	nums05 := [3]int{1, 3, 4}
	nums06 := [3]int{2, 3, 5}

	fmt.Println(nums05 == nums06)
	fmt.Println(nums05 != nums06)

	// 获取数组长度
	fmt.Println(len(nums04))

	// 索引 0，1,2，... len(array) - 1
	fmt.Println(nums04[0], nums04[1])
	nums04[0] = 1000
	fmt.Println(nums04[0])

	for i := 0; i < len(nums04); i++ {
		fmt.Println(nums04[i])
	}

	for _, value := range nums04 { // 只要值不要下标时，可以用空白标识符_去接收它 相当于把它丢弃了
		fmt.Println(value)
	}

	// 切片
	fmt.Printf("%T", nums04[0:15])    // []int 切片类型  不是数组
	fmt.Printf("%T", nums04[1:15:15]) // 切片容量 [start:end:cap]
	fmt.Printf("%v", nums04[1:15:15])

	var marrays [3][2]int // 多维数组  长度为2的int类型的数组 => [2]int   [][] 第一个是它的长度 第二个是它的类型
	fmt.Println(marrays)
	fmt.Println(marrays[0])
	fmt.Println(marrays[0][0])
	marrays[0] = [2]int{1, 3}
	fmt.Println(marrays)
	marrays[1][1] = 1000
	fmt.Println(marrays)

	marrays = [3][2]int{{1, 2}, {3, 4}}
	fmt.Println(marrays)

	for _, v := range marrays {
		for _, vv := range v {
			fmt.Print(vv, "\t")
		}
		fmt.Println()
	}

	// 三维
	var m3 [3][2][5]int
	fmt.Println(m3)

}
