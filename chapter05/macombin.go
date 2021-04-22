package main

import "fmt"

type User struct {
	id   int
	name string
}

func (user User) GetID() int {
	return user.id
}
func (user User) GetName() string {
	return user.name
}
func (user *User) SetID(id int) { // *User 方法接收者要是指针才更新得了值
	user.id = id
}
func (user *User) SetName(name string) {
	user.name = name
}

type Employee struct {
	User
	Salary float64
	name   string
}

func main() {
	var me Employee = Employee{
		User:   User{1, "kk"},
		Salary: 1000,
	}
	fmt.Println(me.User.GetName())
	fmt.Println(me.GetName()) // 可以把匿名的属性去掉
	me.SetName("mian")
	fmt.Println(me.GetName())

}
