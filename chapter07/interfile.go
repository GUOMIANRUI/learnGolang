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
// 谁的方法，是通过接收者定义的，接收者在函数名之前 如(s EmailSender) 就是接收者
func (s EmailSender) Send(to, msg string) error { // s这个变量是EmailSender类型的  只能由EmailSender这个类型去调用Send
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

type WechatSender struct{}

func (s WechatSender) Send(to, msg string) error { // 初始化可以用值也可以指针
	//因为GO里面会为值类型的方法生成指针的
	fmt.Println("发送微信给：", to, "微信内容是：", msg)
	return nil
}

func (s *WechatSender) SendAll(tos []string, msg string) error { // 初始化的时候只能用指针
	for _, to := range tos {
		s.Send(to, msg) // 指针类型的Send方法
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

	sender = &EmailSender{} // 可以传入一个指针类型 值类型会自动生成一个指针类型的接收者
	do(sender)

	sender = &WechatSender{}
	do(sender)
}
