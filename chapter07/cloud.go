package main

import "fmt"

type Host struct {
	ID   string
	Name string
}

type Cloud interface {
	GetList() []Host
	Start(Id string) error
	Stop(Id string) error
	Detail(Id string) Host
}

type Tencent struct {
	API    string
	key    string
	secret string
}

func NewTencent(api, key, secret string) Tencent {
	return Tencent{api, key, secret}
}

func (t Tencent) GetList() []Host {
	fmt.Println("Tencent.GetList")
	return []Host{}
}

func (t Tencent) Start(Id string) error {
	fmt.Println("Tencent.Start")
	return nil
}

func (t Tencent) Stop(Id string) error {
	fmt.Println("Tencent.Stop")
	return nil
}

func (t Tencent) Detail(Id string) Host {
	fmt.Println("Tencent.Detail")
	return Host{}
}

type Aliyun struct {
	API    string
	key    string
	secret string
}

func NewAliyun(api, key, secret string) Aliyun {
	return Aliyun{api, key, secret}
}

func (t Aliyun) GetList() []Host {
	fmt.Println("Aliyun.GetList")
	return []Host{}
}

func (t Aliyun) Start(Id string) error {
	fmt.Println("Aliyun.Start")
	return nil
}

func (t Aliyun) Stop(Id string) error {
	fmt.Println("Aliyun.Stop")
	return nil
}

func (t Aliyun) Detail(Id string) Host {
	fmt.Println("Aliyun.Detail")
	return Host{}
}

func NewCloud(cloudType string) Cloud {
	var cloud Cloud
	if cloudType == "tencent" {
		cloud = NewTencent("tencent", "tencent", "tencent")
	} else if cloudType == "aliyun" {
		cloud = NewAliyun("aliyun", "aliyun", "aliyun")
	} else {
		fmt.Println("error")
	}
	return cloud
}

func main() {
	var cloudType string

	cloudType = "tencent"

	cloud := NewCloud(cloudType)
	if cloud == nil {
		return
	}

	var operate string
EOF:
	for {
		fmt.Print("请输入操作：")
		fmt.Scan(&operate)
		switch operate {
		case "1":
			cloud.GetList()
			fmt.Println("同步主机列表")
			fmt.Println("插入数据库")
			fmt.Println("通知新添加主机")
		case "2":
			cloud.Start("1")
			fmt.Println("启动主机")
		case "3":
			cloud.Stop("2")
			fmt.Println("停止主机")
		case "4":
			cloud.Detail("3")
			fmt.Println("获取主机详情")
		case "5":
			break EOF
		}
	}
}

// 如果一个自己写就完了 如果有两个写if else就完了 但是有3 4 5...个 那就一定要用接口了 如上
