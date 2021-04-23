package main

import "fmt"

// 用结构体的名字去访问一个对象 称为方法表达式
type Dog struct {
	name string
}

func (dog Dog) Call() {
	fmt.Printf("%s: 旺旺\n", dog.name)
}

/*
GO内部自动生成
func (dog *Dog) Call() {
	fmt.Printf("%s: 旺旺\n", dog.name)
}
*/

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func main() {
	m1 := Dog.Call
	m2 := (*Dog).SetName // 定义的时候是指针接收者 所以这里是*Dog

	fmt.Printf("%T, %T\n", m1, m2) // func(main.Dog), func(*main.Dog, string)

	dog := Dog{"豆豆"}
	m1(dog)
	m2(&dog, "小黑") // 方法不会解引用 不能传值类型 得传指针类型
	m1(dog)
	dog.SetName("小白")
	m1(dog)

	pdog := &Dog{"豆豆"}
	m1(*pdog) // *指针 解引用变成值
	m2(pdog, "小黑")
	m1(*pdog)
	m2(pdog, "小白")
	m1(*pdog)
}
