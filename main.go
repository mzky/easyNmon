package main

import (
	"easyNmon/common"
)

func main() {
	//common.InitFlag()
	//common.InitLogs()
	//routers.InitRouter()

	common.Run(1, 100)
	select {}
}
