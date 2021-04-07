# 以gomod的方式管理代码

第一步：初始化 终端输入 go mod init github.com/GUOMIANRUI/usermanager

包的可见性 --> 要在包外可见注意函数名要以大写字母开头 变量也一样

导入一个本地包( --> 模块名+目录名 <-- ) 自己模块下的包 import "github.com/usr/module/pkgname"   第三方的模块 import "github.com/xxxx/xxx"