package main

import (
	"fmt"
	"visibility/users"
)

func main() {
	var u users.User
	// vat a users.address // 小写在包外无法访问

	fmt.Printf("%#v\n", u)

	fmt.Println(u.ID)
	fmt.Println(u.Name)
	// fmt.Println(u.birthday)
	// fmt.Println(u.addr) // 小写的属性在包外也是不能访问的
}
