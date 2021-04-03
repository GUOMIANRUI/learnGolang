package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	maxAuth  = 3
	password = "123456"
)

/*
	命令行的用户管理
	用户信息存储
		-> 内存
		-> 结构 [] map
		-> 用户 id name age tel addr -> 使用映射map  方便通过key查找值 值类型使用string
	用户添加
	用户修改
		请输入需要修改的用户ID
		users[id] --> 在 不在（你输入的用户ID不正确）
		打印用户信息，提示用户是否确认修改（Y/N）
		name, age, tel
	用户删除
		请输入需要修改的用户ID
		users[id] --> 在 不在（你输入的用户ID不正确）
		打印用户信息，提示用户是否确认删除（Y/N）
		Y delete
	用户的查询

	用户登录：
	在程序中定义PASSWORD = "123adc!@#"
	提示输入密码，如果密码输出3次都失败，提示并退出
	如果密码成功，再进行用户管理操作
*/

func inputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

func auth() bool {
	var input string
	for i := 0; i < maxAuth; i++ {
		fmt.Print("请输入密码：")
		fmt.Scan(&input)
		if password == input {
			return true
		} else {
			fmt.Println("密码错误")
		}
	}
	return false
}

func inputUser() map[string]string {
	user := map[string]string{}
	user["name"] = inputString("请输入名字：")
	user["age"] = inputString("请输入年龄：")
	user["tel"] = inputString("请输入联系方式：")
	user["addr"] = inputString("请输入联系地址：")
	return user
}

func getId(users map[int]map[string]string) int { // 获取users里面最大的id +1给将添加的用户
	var id int
	for k, _ := range users {
		if id < k {
			id = k // 获取key的最大值
		}
	}
	return id + 1
}

// 添加用户
// func add(pk int, users map[int]map[string]string) {
func add(users map[int]map[string]string) {
	id := getId(users)
	user := inputUser()
	users[id] = user
	fmt.Println("添加成功")
	// 	var (
	// 		id   string = strconv.Itoa(pk) // 或 fmt.Sprintf("%d", ok) 将int变成str
	// 		name string
	// 		age  string
	// 		tel  string
	// 		addr string
	// 	)
	// 	fmt.Print("请输入姓名：")
	// 	fmt.Scan(&name)

	// 	fmt.Print("请输入年龄：")
	// 	fmt.Scan(&age)

	// 	fmt.Print("请输入联系方式：")
	// 	fmt.Scan(&tel)

	// 	fmt.Print("请输入家庭地址：")
	// 	fmt.Scan(&addr)

	// 	users[id] = map[string]string{
	// 		"id":   id,
	// 		"name": name,
	// 		"age":  age,
	// 		"tel":  tel,
	// 		"addr": addr,
	// 	}
}

func printUser(pk int, user map[string]string) {
	fmt.Println("ID:", pk)
	fmt.Println("名字:", user["name"])
	fmt.Println("年龄:", user["age"])
	fmt.Println("联系方式:", user["tel"])
	fmt.Println("联系地址:", user["addr"])
}

// 查询用户
// q = "" 显示全部
// 非空， 名称 电话 住址 任意一个属性中包含q内容的显示
func query(users map[int]map[string]string) {
	// var q string
	// fmt.Print("请输入查询信息：")
	// fmt.Scan(&q)
	q := inputString("请输入查询信息：")

	title := fmt.Sprintf("%5s|%20s|%5s|%20s|%50s", "ID", "Name", "Age", "Tel", "Addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
	// 注意  下面遍历出来的是 k --> [int]  v --> [string]string    比如 k 为 1   v 为{"name":"xiaoming","age":"20"}
	for k, v := range users {
		if strings.Contains(v["name"], q) || strings.Contains(v["tel"], q) || strings.Contains(v["addr"], q) {
			fmt.Printf("%5d|%20s|%5s|%20s|%50s", k, v["name"], v["age"], v["tel"], v["addr"])
			fmt.Println()
			fmt.Println(strings.Repeat("-", len(title)))
		}
	}
}

// 修改
func modify(users map[int]map[string]string) {
	idString := inputString("请输入修改用户ID：")
	if id, err := strconv.Atoi(idString); err == nil { // 字符串转换为int
		if user, ok := users[id]; ok { // 判断一个key是否在一个映射里，赋值 用第二个值判断
			fmt.Println("将修改的用户信息：")
			// fmt.Println(user)
			printUser(id, user)
			input := inputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				user := inputUser()
				users[id] = user
				fmt.Println("修改成功")
			}
		} else {
			fmt.Println("用户ID不存在")
		}
	} else {
		fmt.Println("输入ID不正确")
	}
}

func del(users map[int]map[string]string) {
	idString := inputString("请输入删除用户ID：")
	if id, err := strconv.Atoi(idString); err == nil { // 字符串转换为int
		if user, ok := users[id]; ok {
			fmt.Println("将删除的用户信息：")
			// fmt.Println(user)
			printUser(id, user)
			input := inputString("确定删除(Y/N)?")
			if input == "y" || input == "Y" {
				delete(users, id)
				fmt.Println("删除成功")
			}
		} else {
			fmt.Println("用户ID不存在")
		}
	} else {
		fmt.Println("输入ID不正确")
	}
}

func main() {

	// 存储用户信息
	if !auth() {
		fmt.Printf("密码%d次错误，程序退出\n", maxAuth)
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
