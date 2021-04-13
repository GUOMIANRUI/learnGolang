package main

import "fmt"

type Address struct {
	Region string
	Street string
	Number string
}

type User struct {
	ID   int
	Name string
	Addr Address
}

type Employee struct {
	*User  // 指针类型匿名结构体
	Salary float64
	Name   string
}

func main() {
	var me Employee
	fmt.Printf("%#v\n", me)

	me02 := Employee{
		User: &User{
			ID:   1,
			Name: "LL",
			Addr: Address{"襄阳市", "江苏路", "002"},
		},
		Salary: 10000,
		Name:   "DD",
	}

	fmt.Printf("%#v\n", me02)
	fmt.Println(me02.Name)
	fmt.Println(me02.Addr)

	fmt.Println(me02.User.Name)
	fmt.Println(me02.User.Addr)

}
