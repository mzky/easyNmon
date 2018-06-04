package main
import (
	"github.com/gin-gonic/gin"
	"flag"
	"os"
	"net/http"
	"os/exec"
	"time"
	"fmt"
	"net"
)
func main() {
	ip := ""
	netaddr, _ := net.InterfaceAddrs()
        networkIp, _ := netaddr[1].(*net.IPNet)
        if !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
        	ip = networkIp.IP.String()
        }	

	port := flag.String("port","","默认监听端口8080\r\n\t设置端口示例：./monitor -port 9999\r\n")
	flag.String("启动监控","","参数n的值：name 生成报告的文件名\r\n\t参数t的值：time 监控时长，单位分钟\r\n\tget_url示例：http://"+ip+":8080/start?n=test&t=30\r\n")
	flag.String("停止所有监控任务","","等同于pkill nmon\r\n\tget_url示例：http://"+ip+":8080/stop\r\n")
	flag.String("查看报告","","浏览器访问url：http://"+ip+":8080/report\r\n")
	flag.String("退出程序","","等同于pkill monitor\r\n\tget_url示例：http://"+ip+":8080/close\r\n")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//浏览报告
	r.StaticFS("/report",http.Dir("./report/"))

	r.GET("/start", func(c *gin.Context) {//格式 ?n=name&t=time 其中&后可为空 默认30分钟
		name := c.DefaultQuery("n", "name")  //取name值
		time := c.DefaultQuery("t", "30") //取执行时间,单位分钟
		lsCmd := exec.Command("/bin/sh", "-c", "./nmonCTL.sh "+name+" "+time)
		
		go func(){
			err := lsCmd.Run()  
			if err!=nil{
			       	fmt.Println(err)
			}	
		}()
		c.JSON(200, gin.H{
		      	"message": string("已执行"+name+"场景监控，持续时间"+time+"分钟"),
		})
	})
	r.GET("/close", func(c *gin.Context) {
		c.JSON(200, gin.H{
		      	"message": "结束程序!",
		})
		go func() {
			time.Sleep(time.Second * 2)
			os.Exit(0)
    		}()
	})
	r.GET("/stop", func(c *gin.Context) {
		lsCmd := exec.Command("/bin/sh", "-c", "ps -ef|grep nmon|grep -v grep|awk {'print $2'}|xargs kill -9")
		err := lsCmd.Start()  
		if err!=nil{
		       	fmt.Println(err)
		}	
		c.JSON(200, gin.H{
		      	"message": "已结束所有监听任务!",
		})
	})
	sport := ":"
	sport += *port
	if *port==""{
		sport +="8080"
	}
	r.Run(sport) // listen and serve on 0.0.0.0:8080
}
