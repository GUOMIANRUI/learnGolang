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

func (employee Employee) GetName() string {
	return employee.name
}

func (employee *Employee) SetName(name string) {
	employee.name = name
}

func main() {
	var me Employee = Employee{
		User:   User{1, "kk"},
		Salary: 1000,
		name:   "小小",
	}
	fmt.Println(me.User.GetName()) // kk or 小小 -> kk
	fmt.Println(me.GetName())      // kk or 小小 -> 小小
	me.SetName("mian")             // employee.name or employee.user.name
	// SetName调用的是 employee.name
	fmt.Println(me.GetName()) // "mian" or 小小 -> mian

}
