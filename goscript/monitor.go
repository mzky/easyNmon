package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	ip := ""
	netaddr, _ := net.InterfaceAddrs()
	networkIp, _ := netaddr[1].(*net.IPNet)
	if !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
		ip = networkIp.IP.String()
	}

	port := flag.String("p", "", "默认监听端口8080,自定义端口加 -p 端口号\r\n\t设置端口示例：./monitor -p 9999\r\n")
	flag.String("启动监控", "", "参数n的值：name 生成报告的文件名\r\n\t参数t的值：time 监控时长，单位分钟\r\n\tget_url示例：http://"+ip+":8080/start?n=test&t=30\r\n")
	flag.String("停止所有监控任务", "", "等同于kill掉nmon进程\r\n\tget_url示例：http://"+ip+":8080/stop\r\n")
	flag.String("查看报告", "", "浏览器访问url：http://"+ip+":8080/report\r\n")
	flag.String("退出程序", "", "关闭自身，结束monitor进程\r\n\tget_url示例：http://"+ip+":8080/close\r\n")
	flag.Parse()

	sport := ":"
	if *port == "" {
		sport += "8080"
	} else {
		sport += *port
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//浏览报告
	r.StaticFS("/report", http.Dir("./report/"))

	r.GET("/start", func(c *gin.Context) { //格式 ?n=name&t=time 其中&后可为空 默认30分钟
		name := c.DefaultQuery("n", "name")  //取name值
		timeStr := c.DefaultQuery("t", "30") //取执行时间,单位分钟

		filename := name + time.Now().Format("20060102150405")

		go func() {
			exec.Command("/bin/bash", "-c", "cp -rf templet report/"+filename).Run()
			exec.Command("/bin/bash", "-c", "./nmon -f -t -s "+timeStr+" -c 60 -m report/"+filename+" -F "+name+".csv").Run()
			t, _ := strconv.Atoi(timeStr)
			time.Sleep(time.Second * time.Duration(t*60+2))
			exec.Command("/bin/bash", "-c", "cd report/"+filename+" &&./toHtml.sh "+name+".csv").Run()
		}()
		c.JSON(200, gin.H{
			"状态": string("已执行" + name + "场景监控，持续时间" + timeStr + "分钟"),
			"报告": string("http://" + ip + sport + "/report"),
		})
	})
	r.GET("/close", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "结束监听服务!",
		})
		go func() {
			exec.Command("/bin/bash", "-c", "cd report/&&for i in `ls`;do cd $PWD/$i;if [ \"`ls index.html`\" != \"index.html\" ];then ./toHtml.sh *.csv; fi;cd ..;done").Run()
			time.Sleep(time.Second * 2)
			os.Exit(0)
		}()
	})
	r.GET("/stop", func(c *gin.Context) {
		exec.Command("/bin/bash", "-c", "ps -ef|grep nmon|grep -v grep|awk {'print $2'}|xargs kill -9").Start()
		exec.Command("/bin/bash", "-c", "cd report/&&for i in `ls`;do cd $PWD/$i;if [ \"`ls index.html`\" != \"index.html\" ];then ./toHtml.sh *.csv; fi;cd ..;done").Run()
		c.JSON(200, gin.H{
			"message": "已结束所有nmon任务!",
		})
	})
	r.Run(sport) // listen and serve on 0.0.0.0:8080
}
