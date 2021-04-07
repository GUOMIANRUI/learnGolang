package main

import (
	"fmt"

	"github.com/GUOMIANRUI/usermanager/users" // 导入本地包： 模块名github.com/GUOMIANRUI/usermanager 目录users
)

// users是一个可执行的包 所以要写一个main包

func main() {

	// 存储用户信息
	if !users.Auth() {
		fmt.Printf("密码%d次错误，程序退出\n", users.MaxAuth)
		return
	}
	fmt.Println("欢迎使用本用户管理系统")
	users := make(map[int]map[string]string)
END:
	for {
		// var op string
		fmt.Print(`
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出
`)
		// fmt.Scan(&op)
		// switch op {
		switch inputString("请输入指令：") {
		case "1":
			fmt.Println("您输入的指令是 新建用户")
			// id++
			add(users)
		case "2":
			fmt.Println("您输入的指令是 修改用户")
			modify(users)
		case "3":
			fmt.Println("您输入的指令是 删除用户")
			del(users)
		case "4":
			fmt.Println("您输入的指令是 查询用户")
			query(users)
		case "5":
			fmt.Println("您输入的指令是 退出")
			break END // 利用标签跳出for循环
		default:
			fmt.Println("指令错误，请重新输入")
		}
		// if op == "5" {
		// 	break
		// }
	}
	// 或定义一个映射建立指令与函数的关系
	// callbacks := map[string]func(map[int]map[string]string){
	// 	"1": add,
	// 	"2": modify,
	// 	"3": del,
	// 	"4": query,
	// }
	// for {
	// 	op := inputString("请输入指令：")
	// 	callback, ok := callbacks[op]

	// 	if ok {
	// 		callback(users)
	// 	} else if op == "5" {
	// 		break
	// 	} else {
	// 		fmt.Println("指令错误，请重新输入")
	// 	}
	// }
}
