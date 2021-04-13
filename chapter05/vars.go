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

func change(u User) {
	u.Name = "HHHH"
}

func changePoint(u *User) {
	u.Name = "YYYY"
}

func main() {
	me := User{}
	// 怎么判定一个值是值类型还是引用类型
	// 修改值 赋值给另外一个值 然后看me2和me是不是一样的
	me2 := me
	me2.Name = "ZZ"
	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", me2)
	// main.User{ID:0, Name:"", Addr:main.Address{Region:"", Street:"", Number:""}}
	// main.User{ID:0, Name:"ZZ", Addr:mai....
	// 所以User是 值类型 ！(对me2修改 对me没有影响) 值类型其实就是拷贝

	change(me2)
	// 改的是里面的u的Name 对外面是没有影响的
	fmt.Printf("%#v\n", me2)

	changePoint(&me2)
	// 这时候me2会变 因为传的是地址
	fmt.Printf("%#v\n", me2)

}
