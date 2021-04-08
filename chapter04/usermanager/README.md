# 以gomod的方式管理代码

第一步：初始化 终端输入 go mod init github.com/GUOMIANRUI/usermanager

包的可见性 --> 要在包外可见注意函数名要以大写字母开头 变量也一样

导入一个本地包( --> 模块名+目录名 <-- ) 自己模块下的包 import "github.com/usr/module/pkgname"   第三方的模块 import "github.com/xxxx/xxx" eg:https://pkg.go.dev/github.com/howeyc/gopass 
import "github.com/howeyc/gopass"

常用：
replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // go mod tidy 整理一下 会自动把版本号拉下来
	golang.org/x/mod => github.com/golang/mod v0.4.2
	golang.org/x/net => github.com/golang/net v0.0.0-20210405180319-a5a99cb37ef4
	golang.org/x/sync => github.com/golang/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys => github.com/golang/sys v0.0.0-20210403161142-5e06dd20ab57
	golang.org/x/term => github.com/golang/term v0.0.0-20210406210042-72f3dc4e9b72
	golang.org/x/text => github.com/golang/text v0.3.6
	golang.org/x/tools => github.com/golang/tools v0.1.0
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

用到字典的话记得初始化
eg: var users map[int]map[string]string = make(map[int]map[string]string) 
也可以先定义 然后在init中初始化
var users map[int]map[string]string
func init() {
    users = make(map[int]map[string]string)
}