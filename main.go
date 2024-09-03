package main

import (
	"easyNmon/common"
)

func main() {
	var f common.Flag
	f.InitFlag()
	f.InitLogs()
	f.InitRouter()

	common.Run(1, 100)
	select {}
}
