package main

import (
	"fmt"
	"strings"
)

func main() {
	// https://golang.google.cn/pkg/strings/#TrimSpace
	fmt.Println(strings.Compare("abc", "bac")) //  0 --> a==b, -1 --> a < b,  +1 --> a > b.
	fmt.Println(strings.Contains("abc", "ad")) //  b是否被包含在b中   false
	fmt.Println(strings.ContainsAny("abcddvdvdv", "ad"))
	fmt.Println(strings.Count("abcddvdvdv", "ad"))             // b在a中出现次数  0
	fmt.Println(strings.Fields("abc def\neeee\fddgb\rvdvdvs")) // 按照空白字符分割  [abc def eeee ddgb vdvdvs]
	fmt.Println(strings.HasPrefix("abcd", "ab"))               // 是不是以ab开头
	fmt.Println(strings.HasSuffix("abcd", "d"))                // 是不是以d结尾
	fmt.Println(strings.Index("defabcdef", "def"))             // 索引位置 0
	fmt.Println(strings.LastIndex("defabcdef", "def"))         // 6

	fmt.Println(strings.Split("abcdef;abc;abc", ";"))                    // 分割 [abcdef abc abc]
	fmt.Println(strings.Join([]string{"first", "second", "third"}, ":")) // 连接

	fmt.Println(strings.Repeat("six", 6))                                              // 重复多少次
	fmt.Println(strings.Replace("sometime in my tears i drown", "sometime", "***", 1)) // 替换函数
	fmt.Println(strings.Replace("sometime in my tears i drown", "i", "*", -1))         // -1 替换所有
	fmt.Println(strings.ReplaceAll("sometime in my tears i drown", "i", "*"))

	fmt.Println(strings.ToLower("abcABC")) // 大写变小写
	fmt.Println(strings.ToUpper("abcABC")) // 小写变大写

	fmt.Println(strings.Trim("xyzabcyzx", "xz"))                              // 删除前后匹配的 yzabcy
	fmt.Println(strings.TrimSpace("   sometime in my tears i drown   \n \r")) // 去除头尾的空白字符

}
