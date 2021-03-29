package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix()) // time 的现在的Unix时间戳

	//生成随机数0 - 100
	guess := rand.Intn(100)
	const maxGuessNum = 5 // 定义一个常量
	isOk := false

	for i := 0; i < maxGuessNum; i++ {
		var input int
		fmt.Print("请输入你猜的值：")
		fmt.Scan(&input)

		if guess == input {
			fmt.Printf("太聪明了，你猜了%d次就猜对了", i+1)
			isOk = true
			break
		} else if input > guess {
			fmt.Printf("太大了，你还有%d次机会\n", maxGuessNum-i-1)
		} else {
			fmt.Printf("太小了，你还有%d次机会\n", maxGuessNum-i-1)
		}
	}
	if !isOk {
		fmt.Println("太笨了， 游戏结束")
	}

}
