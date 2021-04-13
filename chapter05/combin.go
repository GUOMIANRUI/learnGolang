package main

import "fmt"

// 结构体组合 嵌入 --> 匿名嵌入
type Address struct {
	Region string
	Street string
	Number string
}

type User struct {
	ID   int
	Name string
	Addr Address // 结构体组合
}

func main() {
	var me01 User
	fmt.Printf("%#v\n", me01)

	addr := Address{"汕头市", "光明大道", "01"}
	me02 := User{
		ID:   1,
		Name: "mian",
		Addr: addr,
	}
	fmt.Printf("%#v\n", me02)

	// 优化 addr只用到一次 可以在定义变量的时候直接创建这么一个对象
	me03 := User{
		ID:   2,
		Name: "rui",
		Addr: Address{
			Region: "襄阳市",
			Street: "武汉路",
			Number: "01",
		}, // 记得加逗号
	}
	fmt.Printf("%#v\n", me03)
	// 访问并修改属性
	fmt.Println(me03.Addr.Region)
	me03.Addr.Region = "十堰市"
	fmt.Println(me03.Addr.Region)
}
