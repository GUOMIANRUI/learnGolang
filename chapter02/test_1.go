package main

import "fmt"

func main() {
	users := []string{"小明", "小红", "小黑", "小明", "小黑", "小红", "小明"}

	// 自己写的
	// var (
	// 	x = 0
	// 	y = 0
	// 	z = 0
	// )
	// for i := 0; i < len(users); i++ {
	// 	if users[i] == "小明" {
	// 		x += 1
	// 	} else if users[i] == "小红" {
	// 		y += 1
	// 	} else {
	// 		z += 1
	// 	}
	// }
	// fmt.Printf("小明：%d  小红：%d  小黑：%d", x, y, z)  // 小明：3  小红：2  小黑：2

	// 参考
	users_stat := make(map[string]int)

	// for _, user := range users {
	// 	if _, ok := users_stat[user]; !ok {
	// 		users_stat[user] = 1
	// 	} else {
	// 		users_stat[user]++
	// 	}
	// }
	// fmt.Println(users_stat) // map[小明:3 小红:2 小黑:2]

	// 简化   go里面映射中如果没有元素的话会返回0  返回0直接++就行了  不用判断
	for _, user := range users {
		users_stat[user]++
	}
	fmt.Println(users_stat) // map[小明:3 小红:2 小黑:2]

}
