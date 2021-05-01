package main

import "fmt"

type Sender interface { // 在接口中主要定义的是它有哪些方法  对行为的抽象
	// 面向对象里面的多态，一个类型 根据你传进来的类型的不同，它的行为是不一样的
	Send(to string, msg string) error
	SendAll(tos []string, msg string) error
}

// 定义一个结构体
type EmailSender struct {
}

// 这个结构体再定义两个方法
func (s EmailSender) Send(to, msg string) error {
	fmt.Println("发送邮件给：", to, "消息内容是：", msg)
	return nil
}

func (s EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

type SmsSender struct{}

func (s SmsSender) Send(to, msg string) error {
	fmt.Println("发送短信给：", to, "短信内容是：", msg)
	return nil
}

func (s SmsSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

func do(sender Sender) {
	sender.Send("领导", "工作日志")
}

func main() {
	// var sender Sender = EmailSender{} // sender是一个接口，但类型是EmailSender结构体
	var sender Sender = SmsSender{} // 后来领导说换成短信直接将上一条注释，加上这一条就好了
	fmt.Printf("%T %v\n", sender, sender)

	sender.Send("mianrui", "早上好")
	sender.SendAll([]string{"牧眠", "牧野"}, "中午好")

	do(sender)

	sender = EmailSender{}
	do(sender)
}
