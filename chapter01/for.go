package main

import "fmt"

func main() {
	// 索引i => 记录已经加到的n
	// 记录结果
	result := 0

	// 初始化子语句； 条件子语句； 后置子语句
	for i := 1; i <= 100; i++ {
		result += i
	}
	fmt.Println(result)

	result = 0
	i := 1
	for i <= 100 {
		result += i
		i++
	}
	fmt.Println(result)

	// 写死循环
	// i = 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 字符串， 数组， 切片， 映射， 管道
	desc := "我爱祖国"
	for i, ch := range desc {
		fmt.Printf("%d %T %q\n", i, ch, ch)
	}

	for i := 0; i <= 5; i++ {
		if i == 3 {
			break // 跳出循环
		}
		fmt.Println(i)
	}
	for j := 0; j <= 5; j++ {
		if j == 3 {
			continue // 跳出本次循环
		}
		fmt.Println(j)
	}
}
