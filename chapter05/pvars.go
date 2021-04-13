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
	Addr *Address // 加了指针
}

func change(u User) {
	u.Name = "HHHH"
}

func changePoint(u *User) {
	u.Name = "YYYY"
}

// 其他编程语言里面都有一个构造函数的 但是GO没有 可以自己new一个 后面接上结构体名字
func NewUser(id int, name string, region, street, number string) User {
	return User{
		// return &User{ 传出来一个指针类型的
		ID:   id,
		Name: name,
		Addr: &Address{region, street, number},
	}
}

func main() {
	// me := User{}
	// me里面的Addr是指针 零值是nil 访问nil的属性会报错
	// 所以得先初始化一下
	me := User{
		ID:   04,
		Name: "WW",
		Addr: &Address{"襄阳市", "江苏路", "123"},
	}
	me2 := me
	me2.Name = "ZZ"
	me2.Addr.Street = "深圳大道"

	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", me2)
	fmt.Println(me.Addr.Street)
	fmt.Println(me2.Addr.Street)
	// 都是深圳大道 因为me2把地址也复制过来 指向地址是一样的

	// 可以给自己的结构体写一个new函数 然后利用new函数创建对象
	ff := NewUser(2, "FF", "北京市", "海淀区", "456")
	fmt.Printf("%#v\n", ff)
}
