package main

import "fmt"

type Dog struct {
	Name string
}

func (dog Dog) Call() {
	fmt.Printf("%s: hello man\n", dog.Name)
}

func (dog Dog) SetName(name string) {
	dog.Name = name
}

func (dog *Dog) PsetName(name string) {
	dog.Name = name
}

func main() {
	// 方法的定义及调用
	dog := Dog{"豆豆"} // 声明一个对象
	dog.Call()       // 由对象来调用

	// 修改dog的名字
	// dog.Name = "小黑"
	// dog.Call()

	// 如果在包外调用 得提供一个方法修改
	dog.SetName("小白") // 不行 结构体是值类型 调用方法也只是复制一份 所以要用指针
	dog.Call()

	(*&dog).PsetName("旺财")
	dog.Call()

	pdog := &Dog{"小黑"} // 指针类型的对象
	(*pdog).Call()     // 把指针变成值*pdog
	// 因为是指针类型 PsetName方法也是指针接收者 可直接调用
	pdog.PsetName("阿黄")
	(*pdog).Call()
}
