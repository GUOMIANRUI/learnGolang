package main

import "fmt"

// 匿名嵌入

type Address struct {
	Region string
	Street string
	Number string
}

type User struct {
	ID   int
	Name string
	Addr *Address // 指针类型的
}

func main() {
	var me01 User
	fmt.Printf("%#v\n", me01)
	// main.User{ID:0, Name:"", Addr:(*main.Address)(nil)} 零值是nil

	// 初始化
	me02 := User{
		ID:   1,
		Name: "kk",
		Addr: &Address{"西安市", "大马路", "01"}, // 注意逗号
	}
	fmt.Printf("%#v\n", me02)
	// 访问 修改 Address里属性
	fmt.Println(me02.Addr.Region)
	me02.Addr.Region = "汕头市"
	fmt.Println(me02.Addr.Region)
}
