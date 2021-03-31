package main

import "fmt"

func main() {
	var scores map[string]int // nil映射  key是string  value是int

	fmt.Printf("%T %#v\n", scores, scores)
	fmt.Println(scores == nil)

	// 字面量

	// scores = map[string]int{}
	scores = map[string]int{"小明": 8, "小王": 9, "小朱": 10}
	fmt.Println(scores)

	// 还可以通过make函数初始化
	scores = make(map[string]int)
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

	scores["小明"] = 9
	fmt.Println(scores)
	scores["小黑"] = 7
	fmt.Println(scores) // key不存在就是添加

	// 删除  delete()
	delete(scores, "小黑")
	fmt.Println(scores)

	fmt.Println(len(scores))
	for k, v := range scores {
		fmt.Println(k, v)
	} // 映射是无序的

	// key 至少可以有 == != 运算，  bool, 整数， 字符串， 数组 都可以作为key
	// value => 任意类型 slice map
	// 名字 => 映射[字符串]字符串{"地方","联系方式", "成绩"}
	var users map[string]map[string]string // 值里面也是一个映射
	users = map[string]map[string]string{"郭棉锐": {"地方": "汕头", "联系方式": "137********", "成绩": "80"}}
	fmt.Printf("%T, %#v\n", users, users)

	_, ok := users["郭牧眠"]
	fmt.Println(ok)
	users["郭牧眠"] = map[string]string{"地方": "牛客"} // 它是一个映射，所以要用映射的字面量对它进行初始化
	users["郭牧眠"]["成绩"] = "9"
	fmt.Println(users)
	delete(users["郭棉锐"], "联系方式")
	fmt.Println(users)

}
