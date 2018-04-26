package main
import (
        "github.com/gin-gonic/gin"
	"flag"
        "os/exec"
        "fmt"
)
func main() {
	port := flag.String("port","8080","http listen port")
	flag.Parse()
        r := gin.Default()
        r.GET("/start", func(c *gin.Context) {//格式 ?v=value&m=minute 其中&后可为空 默认30分钟
		value := c.DefaultQuery("v", "value")  //取value值
		minute := c.DefaultQuery("m", "30") //取执行时间,单位分钟
		lsCmd := exec.Command("/bin/sh", "-c", "./nmonCTL.sh "+value+" "+minute)
		err := lsCmd.Start()  
                if err!=nil{
                       	fmt.Println(err)
                }	
                c.JSON(200, gin.H{
                      	"message": string("已执行"+value+"场景监控，持续时间"+minute+"分钟"),
                })
        })
        r.GET("/stop", func(c *gin.Context) {
		lsCmd := exec.Command("/bin/sh", "-c", "ps -ef|grep nmon|grep -v grep|awk {'print $2'}|xargs kill -9")
		err := lsCmd.Run()  
                if err!=nil{
                       	fmt.Println(err)
                }	
                c.JSON(200, gin.H{
                      	"message": "已结束所有任务!",
                })
        })
	sport := ":"
	sport += *port
        r.Run(sport) // listen and serve on 0.0.0.0:8080
}
