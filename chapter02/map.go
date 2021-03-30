package main

import "fmt"

func main() {
	var scores map[string]int // nil映射

	fmt.Printf("%T %#v\n", scores, scores)
	fmt.Println(scores == nil)

	// 字面量

	// scores = map[string]int{}
	scores = map[string]int{"小明": 8, "小王": 9, "小朱": 10}
	fmt.Println(scores)

	// 还可以通过make函数初始化
	// scores = make(map[string]int)
	fmt.Println(scores)

	// 增，删，改，查
	// key
	fmt.Println(scores["小明"])
	fmt.Println(scores["小红"]) // key不存在 会返回value类型的零值

	if v, ok := scores["小明"]; ok {
		fmt.Println(v)
	}
	if v, ok := scores["小朱"]; ok {
		fmt.Println(v)
	}
}
