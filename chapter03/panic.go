package main

import "fmt"

func test() (err error) { // 抓取panic错误信息  使用err返回    通过返回值将错误信息返回 做进一步判断
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	panic("error")
	return
}

func main() {
	// recover 当发生错误的时候做一些中止处理 错误处理    配合defer一起使用
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err) // error xxxxx
	// 	}
	// }()

	// fmt.Println("main start")
	// panic("error xxxxx") // 抛出一个错误信息 后面就不执行了

	// fmt.Println("over")

	err := test()
	fmt.Println(err)
}
