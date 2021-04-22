// 方法的嵌入
package main

import (
	"fmt"
)

type Address struct {
	Region string
	Street string
	Number string
}

func (addr Address) Addr() string {
	return fmt.Sprintf("%s-%s-%s", addr.Region, addr.Street, addr.Number)
}

// func (addr Address) String() string {
// 	return fmt.Sprintf("%s-%s-%s", addr.Region, addr.Street, addr.Number)
// } 换成String 会自动调用String 的方法

type User struct {
	ID   int
	Name string
	Addr Address
}

func (user User) User() string {
	return fmt.Sprintf("[%d]%s: %s", user.ID, user.Name, user.Addr.Addr())
	// 第一个Addr是user的属性，第二个Addr是Address的方法  即方法的嵌入
}

// func (user User) String() string {
// 	return fmt.Sprintf("[%d]%s: %s", user.ID, user.Name, user.Addr)
// }

func main() {
	var u User = User{
		ID:   1,
		Name: "kk",
		Addr: Address{"西安市", "锦业路", "001"},
	}
	fmt.Println(u.User())
	fmt.Println(u.Addr.Addr())
	// fmt.Println(u)   Print会自动调用String 的方法
	// fmt.Println(u.Addr)
}
