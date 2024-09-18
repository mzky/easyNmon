package main

import (
	"easyNmon/common"
	"easyNmon/router"
)

func main() {
	var f common.Flag
	f.InitFlag()
	f.InitLogs()
	router.Flag(f).InitRouter()

	common.Run(1, 100)
	select {}
}
