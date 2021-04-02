package main

import (
	"errors"
	"fmt"
)

// go里面建议大家自己处理错误 并通过返回值返回
// 返回值 怎么定义错误类型 --> error
// 怎么创建错误类型的值 --> errors.New()  或  fmt.Errorf()
// 无错误返回 nil

func division(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println(division(1, 3))
	if v, err := division(1, 0); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	e := fmt.Errorf("Error: %s", "division by zero")
	fmt.Printf("%T, %v\n", e, e) // *errors.errorString, Error: division by zero
}
