package main

import (
	"easyNmon/common"
	"easyNmon/router"
)

func main() {
	var f common.Flag
	f.InitFlag()
	f.InitLogs()
	router.InitRouter(f)

	common.Run(1, 100)
	select {}
}
