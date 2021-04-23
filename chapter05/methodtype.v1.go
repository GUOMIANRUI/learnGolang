package main

import "fmt"

type Dog struct {
	name string
}

func (dog *Dog) Call() { // 指针接收者
	fmt.Printf("%s: 旺旺\n", dog.name)
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func main() {
	dog := Dog{"豆豆"}

	m1 := dog.Call         // dog 取引用 存储的是dog的指针
	fmt.Printf("%T\n", m1) // fun() function类型
	m1()                   // 豆豆: 旺旺

	dog.SetName("小黑")
	dog.Call() // 小黑: 旺旺
	m1()       // 小黑: 旺旺

	pdog := &Dog{"豆豆"}
	m2 := pdog.Call //
	fmt.Printf("%T\n", m2)
	m2()

	pdog.SetName("小黑")
	pdog.Call() // 小黑: 旺旺
	m2()        // 小黑: 旺旺

}
