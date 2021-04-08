package main

import (
	"fmt"

	upkg "github.com/GUOMIANRUI/usermanager/users" // 导入本地包： 模块名github.com/GUOMIANRUI/usermanager 目录users
	// 起别名 防止包名与下面的变量名冲突
)

// users是一个可执行的包 所以要写一个main包

func main() {

	// 存储用户信息
	if !upkg.Auth() {
		fmt.Printf("密码%d次错误，程序退出\n", upkg.MaxAuth)
		return
	}
	fmt.Println("欢迎使用本用户管理系统")
	users := make(map[int]map[string]string)
END:
	for {
		fmt.Print(`
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出
`)
		switch upkg.InputString("请输入指令：") {
		case "1":
			fmt.Println("您输入的指令是 新建用户")
			upkg.Add(users)
		case "2":
			fmt.Println("您输入的指令是 修改用户")
			upkg.Modify(users)
		case "3":
			fmt.Println("您输入的指令是 删除用户")
			upkg.Del(users)
		case "4":
			fmt.Println("您输入的指令是 查询用户")
			upkg.Query(users)
		case "5":
			fmt.Println("您输入的指令是 退出")
			break END // 利用标签跳出for循环
		default:
			fmt.Println("指令错误，请重新输入")
		}
	}
}
