package main

import "fmt"

type Sender interface {
	Send(msg string) error
}

type Reciver interface {
	Recive() (string, error)
}

type Client interface {
	Sender
	Reciver // 匿名嵌入 Client相当于嵌入了两个方法 Sender Reciver

	Open() error
	Close() error
}

type MSNClient struct{}

func (c MSNClient) Open() error {
	fmt.Println("Open")
	return nil
}

func (c MSNClient) Close() error {
	fmt.Println("Close")
	return nil
}

func (c MSNClient) Send(msg string) error {
	fmt.Println("send:", msg)
	return nil
}

func (c MSNClient) Recive() (string, error) { // 不要忘记逗号
	fmt.Println("recive")
	return "", nil
}

func main() {
	msn := MSNClient{}

	var s Sender = msn
	var r Reciver = msn
	var c Client = msn

	s.Send("1")
	r.Recive()
	c.Send("2")
	c.Recive()

	c.Open()
	defer c.Close() // 延时关闭
	c.Send("3")
	c.Recive()

}
