module github.com/GUOMIANRUI/usermanager

go 1.12

require (
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550 // indirect
)

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
