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
	Addr Address // 结构体嵌入
}

type Employee struct {
	// User User // 定义一个User 他的类型就是用户
	User   // 匿名嵌入  User:main.User
	Salary float64
	Name   string
}

func main() {
	var me Employee
	fmt.Printf("%T, %#v\n", me, me)
	// 初始化
	me02 := Employee{
		User: User{
			ID:   1,
			Name: "mian",
			Addr: Address{"汕头", "大马路", "02"},
		},
		Salary: 10000,
	}
	fmt.Printf("%T, %#v\n", me02, me02)
	fmt.Println(me02.User.Addr.Street)
	me02.User.Addr.Street = "小公路"
	fmt.Printf("%T, %#v\n", me02, me02)

	fmt.Println(me02.Name) // 访问匿名结构体里的属性 可省略匿名属性名User
	fmt.Println(me02.Addr.Street)
	me02.Addr.Street = "二路"
	fmt.Printf("%T, %#v\n", me02, me02)

}
