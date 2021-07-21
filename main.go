package main

import (
	"easyNmon/common"
	"easyNmon/routers"
	"easyNmon/utils"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	//ip := controllers.GetExternalIP()
	ip := ""
	debug := flag.Bool("debug", false, "debug mode")
	version := flag.Bool("v", false, "show version")
	port := flag.String("p", "9999",
		"default port [9999]\ncustom port e.g.：./easyNmon -p 9999")
	dir := flag.String("d", "report",
		"default directory [report]\ncustom directory e.g.：./easyNmon -d /mnt/report")
	analysis := flag.String("a", "",
		"create analysis directory\ne.g.：./easyNmon -a ./report/nmon_testName")
	read := flag.Bool("r", false, "read interface")

	host := fmt.Sprintf("http://%s:%s", ip, *port)
	readme := "web管理\t" + host + "\n" +
		"interface(Get)：\n" +
		"start\t" + host + "/start?n=name&t=30&f=30\n\t" +
		"参数n为生成报告的文件名,\n\t" +
		"参数t为监控时长(单位分钟),\n\t" +
		"参数f为监控频率，每隔多少秒收集一次\n" +
		"stop\t" + host + "/stop\n\t" +
		"停止所有监控任务\n" +
		"report\t" + host + "/report\n\t" +
		"查看报告\n" +
		"close\t" + host + "/close\n\t" +
		"结束EasyNmon进程\n"
	flag.Parse()

	common.InitSetup()

	if *version {
		fmt.Println("Version: " + common.Version)
		fmt.Println("Compile: " + common.Compile)
		os.Exit(0)
	}
	if *read {
		fmt.Println(readme)
		os.Exit(0)
	}
	anal := *analysis
	if anal != "" {
		path, name := filepath.Split(anal)
		utils.GetNmonReport(path, name)
		os.Exit(0)
	}

	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	}

	common.ReportDir, _ = filepath.Abs(*dir) //绝对路径*dir
	syscall.Umask(0)
	if os.MkdirAll(common.ReportDir, os.ModePerm) != nil {
		logrus.Error("easyNmon启动权限不足或当前目录无权写入!")
		os.Exit(0)
	}

	routers.Router()

	common.R.Run(":" + *port) // listen
	logrus.Error("easyNmon启动失败，查看端口是否已被占用!")
}
