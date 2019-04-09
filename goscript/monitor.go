package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var (
	Version   = "0.1"
	BuildTime = "20180808"
)

func main() {
	ip := ""
	netaddr, _ := net.InterfaceAddrs()
	networkIp, _ := netaddr[1].(*net.IPNet)
	if !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
		ip = networkIp.IP.String()
	}

	version := flag.Bool("v", false, "显示版本号")
	port := flag.String("p", "9999", "默认监听端口9999,自定义端口加 -p 端口号\n设置端口示例：./monitor -p 9999")
	flag.Bool("web管理页面", false, "浏览器访问http://"+ip+":9999")
	flag.Bool("启动监控", false, "参数n的值：name 生成报告的文件名\n参数t的值：time 监控时长，单位分钟\n\tget_url示例：http://"+ip+":9999/start?n=test&t=30")
	flag.Bool("停止所有监控任务", false, "等同于kill掉nmon进程\nget_url示例：http://"+ip+":9999/stop")
	flag.Bool("查看报告", false, "浏览器访问：http://"+ip+":9999/report，也可通过web管理页面入口查看")
	flag.Bool("退出程序", false, "关闭自身，结束monitor进程\nget_url示例：http://"+ip+":9999/close")
	flag.Parse()

	if *version {
		fmt.Println("Version: " + Version)
		fmt.Println("BuildTime: " + BuildTime)
		return
	}

	fmt.Println("访问web管理页面 : http://" + ip + ":" + *port)
	fmt.Println("启动监控接口示例: http://" + ip + ":" + *port + "/start?n=testname&t=30")
	fmt.Println("停止所有监控接口: http://" + ip + ":" + *port + "/stop")
	fmt.Println("浏览器查看报告 :  http://" + ip + ":" + *port + "/report")
	fmt.Println("结束monitor进程:  http://" + ip + ":" + *port + "/close")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//首页
	r.LoadHTMLGlob("web/index.html")
	r.Static("/js", "web/js")
	r.Static("/fonts", "web/fonts")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	os.MkdirAll("./report", 777)
	//浏览报告
	r.StaticFS("/report", http.Dir("report"))
	//r.StaticFS("/", http.Dir("web"))

	r.GET("/start", start)
	r.GET("/close", close)
	r.GET("/stop", stop)

	r.Run(":" + *port) // listen and serve on 0.0.0.0:8080
}

func start(c *gin.Context) { //格式 ?n=name&t=time 其中&后可为空 默认30分钟
	name := c.DefaultQuery("n", "name")  //取name值
	timeStr := c.DefaultQuery("t", "30") //取执行时间,单位分钟

	filename := name + time.Now().Format("20060102150405")

	go func() {
		exec.Command("/bin/bash", "-c", "cp -rf templet report/"+filename).Run()
		exec.Command("/bin/bash", "-c", "./nmon -f -t -s "+timeStr+" -c 60 -m report/"+filename+" -F "+name).Run()
		t, _ := strconv.Atoi(timeStr)
		time.Sleep(time.Second * time.Duration(t*60+2))
		exec.Command("/bin/bash", "-c", "cd report/"+filename+" &&./toHtml.sh "+name).Run()
	}()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": string("已执行" + name + "场景监控，测试时长 " + timeStr + " 分钟"),
	})
}

func close(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "已结束EasyNmon服务!",
	})
	go func() {
		exec.Command("/bin/bash", "-c", "cd report/&&for i in `ls`;do cd $PWD/$i;if [ \"`ls index.html`\" != \"index.html\" ];then ./toHtml.sh \"`ls|grep -v js|grep -v templet|grep -v toHtml.sh`\"; fi;cd ..;done").Run()
		time.Sleep(time.Second * 2)
		os.Exit(0)
	}()
}

func stop(c *gin.Context) {
	exec.Command("/bin/bash", "-c", "ps -ef|grep nmon|grep -v grep|awk {'print $2'}|xargs kill -9").Start()
	exec.Command("/bin/bash", "-c", "cd report/&&for i in `ls`;do cd $PWD/$i;if [ \"`ls index.html`\" != \"index.html\" ];then ./toHtml.sh \"`ls|grep -v js|grep -v templet|grep -v toHtml.sh`\"; fi;cd ..;done").Run()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "已结束所有服务器监控任务!",
	})
}
