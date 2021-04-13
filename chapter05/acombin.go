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

type Company struct {
	ID     int
	Name   string
	Addr   Address
	Salary float64
}

type Employee struct {
	// User User // 定义一个User 他的类型就是用户
	User
	Company
	Salary float64
}

func main() {
	var me Employee
	fmt.Printf("%T, %#v\n", me, me)

	me.Company.Name = "AA"
	me.User.Name = "BB"
	fmt.Println(me.Company.Name) // 两个匿名结构体属性冲突时必须使用全路径的方式访问
	fmt.Println(me.User.Name)

	fmt.Printf("%#v\n", me)
}
