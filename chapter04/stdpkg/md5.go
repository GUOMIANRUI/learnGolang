package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {

	// fmt.Printf("%x", md5.Sum([]byte("i am mianrui"))) // 得到md5值 470e4d59e95c6d1006779e8a76bd321c

	bytes := md5.Sum([]byte("i am mianrui"))
	x := fmt.Sprintf("%x", bytes) // 转换为字符串
	// 或
	fmt.Println(hex.EncodeToString(bytes[:])) // 传入一个切片类型
	fmt.Println(x)

	// 数据比较大的时候要设置一个md5值的对象 可以分多次计算md5值
	m := md5.New()
	m.Write([]byte("i am"))
	m.Write([]byte(" mianrui")) // 470e4d59e95c6d1006779e8a76bd321c

	fmt.Printf("%x\n", m.Sum(nil)) // 传一个空的切片

	// shal的写法也和md5值类似 sha256也是
	// h := shal.New()
	// h := sha256.New()

	key := []byte("1234567890abcdefg")
	// 创建hmac hash 对象
	hash := hmac.New(sha256.New, key)
	// 写入字符串计算散列
	io.WriteString(hash, "Hi, guomianrui")
	// 计算hmac散列
	fmt.Printf("%x\n", hash.Sum(nil))
	// 判断是否相等
	hash2 := hmac.New(sha256.New, key)
	hash2.Write([]byte("Hi, guomianrui"))
	fmt.Println(hmac.Equal(hash2.Sum(nil), hash.Sum(nil)))
}
