package main

import "fmt"

func main() {
	// 切片定义：  []int{} => 空切片   底层的数组没有存放元素
	//            var nums []int => nil 切片    底层数组的指针是指向nil的

	var nums []int

	fmt.Printf("%T\n", nums)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	fmt.Println(nums == nil) // 切片底层指向空

	// 字面量
	nums = []int{1, 2, 3}
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums)) // %#v 显示类型和值
	nums = []int{1, 2, 3, 4}
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))

	// 数组切片赋值
	var arrays [10]int = [10]int{1, 2, 3, 4, 5, 6}
	nums = arrays[1:10]
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))

	// 使用make函数初始化切片
	nums = make([]int, 3)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	nums = make([]int, 3, 5)                              // 长度 容量
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums)) // 容量的设计让其长度可变

	// 元素操作（增删改查）
	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])
	// fmt.Println(nums[3]) // runtime error: index out of range  超出长度的索引不能使用
	// nums[3] = 10

	nums = append(nums, 1) // 添加元素  长度加一
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	nums = append(nums, 1) // 添加元素  长度加一
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	nums = append(nums, 1) // 添加元素  长度加一 容量不够的话会扩展 1 1.5倍
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for index, value := range nums {
		fmt.Println(index, value)
	}

	nums = make([]int, 3, 10)
	// 切片操作
	n := nums[1:3:8]
	// n_cap - start
	fmt.Printf("%T %#v %d %d\n", n, n, len(n), cap(n)) // []int []int{0, 0} 2 7
	n = nums[2:3]
	// src_nap - start
	fmt.Printf("%T %#v %d %d\n", n, n, len(n), cap(n)) // []int []int{0} 1 8

	// 容量产生的影响  两个切片底层会共用一个数组 产生一些副作用   如果其中一个容量不够，重新申请数组，这个影响就会消失
	nums = make([]int, 3, 5)
	nums02 := nums[1:3]
	fmt.Println(nums, nums02)
	nums02[0] = 1
	fmt.Println(nums, nums02)

	// 这个副作用还可以延伸到数组上面
	nums = arrays[:]
	fmt.Println(nums, arrays)

	nums[0] = 100 // arrays[0] 也会改变
	fmt.Println(nums, arrays)

	// 了解即可 以后写数组切片的时候注意上面几点即可

	// 删除 copy
	nums04 := []int{1, 2, 3}
	nums05 := []int{10, 20, 30, 40}
	copy(nums05, nums04)
	fmt.Println(nums05)

	nums05 = []int{10, 20, 30, 40}
	copy(nums04, nums05)
	fmt.Println(nums04)

	// 删除 索引为0  索引最后一个
	nums06 := []int{1, 2, 3, 4, 5}
	fmt.Println(nums06[1:])
	fmt.Println(nums06[:len(nums06)-1])
	// 删除 中间的元素 2(3)
	copy(nums06[2:], nums06[3:])
	fmt.Println(nums06[:len(nums06)-1])

	// 堆栈 先进后出 队尾添加 队尾移除
	// 队列 先进先出 队尾添加 队头移除
	queue := []int{}
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	queue = append(queue, 5)
	//1, 2, 3, 5
	fmt.Println(queue)
	fmt.Print(queue[0])
	queue = queue[1:]
	// 2,3,5
	fmt.Println(queue)
	fmt.Print(queue[0])
	queue = queue[1:]
	// 3,5
	fmt.Println(queue)

	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)

	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[len(stack)-1])

	// 多维切片
	points := [][]int{}
	points02 := make([][]int, 0) // 创建二维切片的两种方法
	fmt.Printf("%T\n", points02)

	points = append(points, []int{1, 2, 3})
	points = append(points, []int{3, 4, 0})
	points = append(points, []int{3, 4, 0, 2, 4, 5})
	fmt.Println(points)
	fmt.Println(points[0])
	fmt.Println(points[0][1])

	// 切片和数组的区别   数组是值类型，有对象指向他时，会复制一块空间存放元素

	slice01 := []int{1, 2, 3}
	slice02 := slice01
	slice02[0] = 10 // slice01会变
	fmt.Println(slice01, slice02)

	array01 := [3]int{1, 2, 3}
	array02 := array01
	array02[0] = 10 //   array01不会变
	fmt.Println(array01, array02)

}
