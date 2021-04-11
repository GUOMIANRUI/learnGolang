package main

import (
	"fmt"
	"time"
)

// 自定义结构体
// 定义一个类型：User   模式是：大括号里(即这个类型都有哪些属性)
type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
	Remark   string
}

func main() {
	// 定义这个User类型的变量怎么定义： var me User
	var me User
	fmt.Printf("%T\n", me)
	fmt.Printf("%#v\n", me) // main.User{ID:0, Name:"", Birthday:
	//time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)},
	// Addr:"", Tel:"", Remark:""}
	// 结构体的零值不是nil 而是每个子元素的零值组成的一个结构
	// ID的零值是0 Name是nil Birthday的Time结构体里面子元素的零值（
	// 看源码Time也是一个结构体)

	// 根据字面量值对它进行初始化 必须每个属性写上值 按顺序定义
	var me2 User = User{1, "me2", time.Now(), "汕头市", "123456", "123"}
	fmt.Printf("%#v\n", me2)

	// 显式地使用零值初始化
	var me3 User = User{} // 同var me3 User 但后面要用到指针的话要用前者
	fmt.Printf("%#v\n", me3)

	// 指定属性名进行初始化 顺序可以不相关 还可以省略一些值 没写的用零值初始化
	var me4 User = User{Name: "someone", ID: 1, Birthday: time.Now(), Addr: "汕头", Tel: "123", Remark: "132"}
	fmt.Printf("%#v\n", me4)

	// 也可以这样写 数组 切片也都可以 但注意最后一个逗号
	var me5 User = User{
		Name:     "someone2",
		ID:       1,
		Birthday: time.Now(),
		Addr:     "汕头",
		Tel:      "456",
	}
	fmt.Printf("%#v\n", me5)

	// 定义一个指针类型
	var pointer *User // 没有初始化 就是零值 nil
	fmt.Printf("%T\n", pointer)
	fmt.Printf("%#v\n", pointer) // (*main.User)(nil)

	// 初始化
	// 通过已有变量取地址定义
	var pointer2 *User = &me4
	fmt.Printf("%#v\n", pointer2) // &main.User{ID:1 ....
	// 创建一个结构体的零值并取地址 &User{} 然后赋值给指针
	var pointer3 *User = &User{}
	fmt.Printf("%#v\n", pointer3)
	// 使用new() 和上面类似 打印结果一样 因为new作用是创建某类型的指针并初始化
	var pointer4 *User = new(User)
	fmt.Printf("%#v\n", pointer4)

	// 包括之前的int float ... 想创建其对应指针类型的零值也可以用new()
	var pointer5 *int = new(int)
	fmt.Printf("%#v\n", pointer5) // (*int)(0xc0000682e0)
	// 但是对于切片和映射一般不会用new()
	// 因为它俩本身是一个引用类型 所以使用make()
}
