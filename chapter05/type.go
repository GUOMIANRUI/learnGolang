package main

import "fmt"

type Counter int // 定义一个类型

type User map[string]string

type Callback func(...string)

func main() {
	var counter Counter = 20

	counter += 10
	fmt.Println(counter)

	me := make(User)
	me["name"] = "kk"
	me["addr"] = "汕头"
	fmt.Println(me)
	fmt.Printf("%T, %T\n", counter, me) // main.Counter, main.User

	var list Callback = func(args ...string) {
		for i, v := range args {
			fmt.Println(i, ":", v)
		}
	}

	list("a", "b", "c")

	var counter2 int = 10
	// 转换为同一类型才能比较
	fmt.Println(int(counter) > counter2)
	fmt.Println(counter > Counter(counter))
}
