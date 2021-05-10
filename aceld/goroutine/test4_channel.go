package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			//如果c可写，则读case就会进来
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return // 如果可读 退出循环
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	//sub go
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}

		quit <- 0 // 会触发上面那个case
	}()

	// main go
	fibonacii(c, quit)
}
