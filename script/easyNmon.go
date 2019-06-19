package main

import (
	"easyNmon/internal"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

var (
	// 版本号
	Version = "<Version>"
	// 构建时间
	BuildTime = "<BuildTime>"
	//全局路径
	ReportDir string
)

func main() {
	ip := ""
	netaddr, _ := net.InterfaceAddrs()
	networkIp, _ := netaddr[1].(*net.IPNet)
	if !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
		ip = networkIp.IP.String()
	}

	version := flag.Bool("v", false, "显示版本号")
	port := flag.String("p", "9999", "默认监听端口9999,自定义端口加 -p 端口号\n设置端口示例：./easyNmon -p 9999")
	dir := flag.String("d", "report", "指定生成报告的directory")
	flag.Bool("web管理页面", false, "浏览器访问http://"+ip+":9999")
	flag.Bool("启动监控", false, "参数n的值：name 生成报告的文件名\n参数t的值：time 监控时长，单位分钟\nget_url示例：http://"+ip+":9999/start?n=test&t=30")
	flag.Bool("停止所有监控任务", false, "等同于kill掉nmon进程\nget_url示例：http://"+ip+":9999/stop")
	flag.Bool("查看报告", false, "浏览器访问：http://"+ip+":9999/report，也可通过web管理页面入口查看")
	flag.Bool("退出程序", false, "关闭自身，结束easyNmon进程\nget_url示例：http://"+ip+":9999/close")
	flag.Parse()

	ReportDir = *dir
	os.MkdirAll(ReportDir, 777)

	if *version {
		fmt.Println("Version: " + Version)
		fmt.Println("BuildTime: " + BuildTime)
		os.Exit(0)
	}

	fmt.Println("访问web管理页面 : http://" + ip + ":" + *port)
	fmt.Println("启动监控接口示例: http://" + ip + ":" + *port + "/start?n=testname&t=30")
	fmt.Println("停止所有监控接口: http://" + ip + ":" + *port + "/stop")
	fmt.Println("浏览器查看报告 :  http://" + ip + ":" + *port + "/report")
	fmt.Println("结束easyNmon进程:  http://" + ip + ":" + *port + "/close")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//重定向首页--解决静态文件与接口共存
	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/web"
		defer r.HandleContext(c)
	})
	//首页
	r.StaticFS("/web", http.Dir("web"))
	// 浏览报告
	r.StaticFS("/report", http.Dir(ReportDir))
	//生成报告,用于实时更新报告
	r.GET("/generate/:name/", func(c *gin.Context) {
		name := c.Param("name")
		internal.GetNmonReport(filepath.Join(ReportDir, name), name[:len(name)-14])
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.GET("/start", start)
	r.GET("/close", close)
	r.GET("/stop", stop)

	r.Run(":" + *port) // listen and serve on 0.0.0.0:8080
}

func start(c *gin.Context) { // 格式 ?n=name&t=time 其中&后可为空 默认30分钟
	name := c.DefaultQuery("n", "name")  // 取name值
	timeStr := c.DefaultQuery("t", "30") // 取执行时间,单位分钟
	filename := name + time.Now().Format("20060102150405")

	go func() {
		fp := filepath.Join(ReportDir, filename)
		os.MkdirAll(fp, 777)
		exec.Command("/bin/bash", "-c", "./nmon -f -t -s "+timeStr+" -c 60 -m "+fp+" -F "+name).Run()
		t, _ := strconv.Atoi(timeStr)
		time.Sleep(time.Second * 2)
		internal.GetNmonReport(fp, name)
		time.Sleep(time.Second * time.Duration(t*60+2))
		internal.GetNmonReport(fp, name)
	}()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": string("已执行" + name + "场景监控，测试时长 " + timeStr + " 分钟"),
	})
}

func close(c *gin.Context) { //结束自身进程
	getAllReport()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "已结束EasyNmon服务!",
	})
	go func() {
		defer os.Exit(0)
	}()
}

func stop(c *gin.Context) {
	getAllReport()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "已结束所有服务器监控任务!",
	})
}

//重新生成所有报告
func getAllReport() {
	list, _ := getDirList(ReportDir)
	for _, v := range list {
		internal.GetNmonReport(filepath.Join(ReportDir, v), v[:len(v)-14])
	}
}

//获取文件夹列表
func getDirList(dirpath string) ([]string, error) {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				if path != dirpath {
					dir_list = append(dir_list, path[len(ReportDir)+1:])
					return nil
				}
			}
			return nil
		})
	return dir_list, dir_err
}
