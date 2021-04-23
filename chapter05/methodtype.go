package main

import "fmt"

type Dog struct {
	name string
}

func (dog Dog) Call() { // 值接收者
	fmt.Printf("%s: 旺旺\n", dog.name)
}

func (dog *Dog) SetName(name string) { // 指针接收者
	dog.name = name
}

func main() {
	dog := Dog{"豆豆"}

	m1 := dog.Call         // 复制一份给m1 所以后面改成小黑对m1没有影响
	fmt.Printf("%T\n", m1) // fun() function类型
	m1()

	dog.SetName("小黑")
	dog.Call() // 小黑: 旺旺
	m1()       // 豆豆: 旺旺

	pdog := &Dog{"豆豆"}
	m2 := pdog.Call // pdog会自动解引用 之后是值类型了 也是复制一份
	fmt.Printf("%T\n", m2)
	m2()

	pdog.SetName("小黑")
	pdog.Call() // 小黑: 旺旺
	m2()        // 豆豆: 旺旺

}
