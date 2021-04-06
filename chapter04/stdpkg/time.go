package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T\n", now)
	fmt.Printf("%v\n", now)

	// 2000/01/02 08:10:01
	// 2006 年
	// 01 月 02 天 24小时制用15表示  分钟04  秒 05
	fmt.Println(now.Format("2006/01/02 15:04:05")) // 2021/04/06 20:39:53 把now的Time类型改为字符串类型
	fmt.Println(now.Format("2006-01-02 15:04:05")) // 2021/04/06 20:39:53 24小时制用15表示
	fmt.Println(now.Format("2006-01-02"))          // 2021-04-06
	fmt.Println(now.Format("15:04:05"))            // 20:39:53

	fmt.Println(now.Unix()) // 1617712793
	fmt.Println(now.UnixNano())

	t, err := time.Parse("2006-01-02 15:04:05", "2006-01-02 03:06:09") // 定义时间的格式 后面的符合前面的格式的话会返回<nil> 不符合返回error
	fmt.Println(err, t)

	t = time.Unix(0, 0)
	fmt.Println(t) // 1970-01-01 08:00:00 +0800 CST

	d := now.Sub(t)            // 可用于计算两个时间的时间区间
	fmt.Printf("%T, %v", d, d) // Duration 时间区间类型

	// 一些常量 GO定义的时间区间    time.Second   time.Minute   time.Hour

	fmt.Println(time.Now())
	time.Sleep(time.Second * 5) // 休眠5秒
	fmt.Println(time.Now())

	k := now.Add(3 * time.Hour) // 3小时后的时间
	fmt.Println(k)

	d, err = time.ParseDuration("3h2m4s")
	fmt.Println(err, d)    // <nil> 3h2m4s
	fmt.Println(d.Hours()) // 它就把时间区间转化为多少个小时 3.0344444444444445
	fmt.Println(d.Minutes())
	fmt.Println(d.Seconds())
}
