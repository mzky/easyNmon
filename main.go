package main

import (
	"easyNmon/common"
	"easyNmon/router"
	"github.com/mzky/utils/log"
)

func main() {
	log.InitLog("debug", "easyNmon.log")
	var flag common.Flag
	common.InitFlag(&flag)

	common.Run("3", "100")

	router.InitRouter(&flag)

	select {}
}
