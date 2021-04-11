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
	var me User = User{ // 或简短声明 me := User
		Name:     "someone2",
		ID:       1,
		Birthday: time.Now(),
		Addr:     "汕头",
		Tel:      "456",
	}
	// 访问变量属性 变量名.属性名
	fmt.Println(me.ID, me.Name, me.Tel)
	// 修改属性值
	me.Tel = "123456789"
	fmt.Printf("%#V\n", me)

	// 指针的结构体变量如何访问和修改
	me2 := &User{
		ID:   2,
		Name: "suibian",
	}
	fmt.Printf("%T\n", me2)

	(*me2).Tel = "137xxxxxxxx" // me2是指针 解引用*me2得到值
	// 在GO里面可以省略 直接
	me2.Addr = "潮阳"
	fmt.Printf("%#V\n", me2)
}
