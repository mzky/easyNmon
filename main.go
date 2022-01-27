package main

import (
	"easyNmon/common"
	"easyNmon/routers"
)

func main() {
	common.InitFlag()
	common.InitLogs()
	routers.InitRouter()
}
