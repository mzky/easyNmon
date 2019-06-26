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
	"strings"
	"syscall"
	"time"
)

var (
	// 版本号
	Version = "<Version>"
	// 构建时间
	BuildTime = "<BuildTime>"
	//全局
	ReportDir string
	NmonPath  string
)

func main() {
	ip := ""
	netaddr, _ := net.InterfaceAddrs()
	networkIp, _ := netaddr[1].(*net.IPNet)
	if !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
		ip = networkIp.IP.String()
	}

	readme := "接口(Get)：\n\t/start\t启动监控,参数n为生成报告的文件名,参数t为监控时长(单位分钟)：\n\t\thttp://" + ip + ":9999/start?n=name&t=30\n\t/stop\t停止所有监控任务：\n\t\thttp://" + ip + ":9999/stop\n\t/report\t查看报告：\n\t\thttp://" + ip + ":9999/report\n\t/close\t关闭自身：\n\t\thttp://" + ip + ":9999/close\n管理页面：\n\t通过浏览器访问web管理页面：\n\thttp://" + ip + ":9999"
	version := flag.Bool("v", false, "version:显示版本号")
	port := flag.String("p", "9999", "port:默认监听端口9999,自定义端口加 -p 端口号\n示例：./easyNmon -p 9999")
	dir := flag.String("d", "report", "directory:指定生成报告的路径\n示例：./easyNmon -d /mnt/rep")
	analysis := flag.String("a", "", "analysis:生成html图表，参数指定nmon报告文件，同目录生成html图表\n示例：./easyNmon -a ./report/nmonTestName")
	nmonpath := flag.String("np", "nmon/nmon", "nmonpath：指定对应系统版本的nmon文件\n示例：./easyNmon -np ./nmon/nmon_xxx")
	flag.Bool("操作说明", false, readme)
	flag.Parse()

	ReportDir = *dir
	NmonPath = *nmonpath
	os.MkdirAll(ReportDir, 777)

	if *version {
		fmt.Println("Version: " + Version)
		fmt.Println("BuildTime: " + BuildTime)
		os.Exit(0)
	}

	analy := *analysis
	if analy != "" {
		paths, fileName := filepath.Split(analy)
		internal.GetNmonReport(paths, fileName)
		os.Exit(0)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//重定向首页--解决静态文件与接口共存
	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/web"
		defer r.HandleContext(c)
	})
	//首页
	r.StaticFS("/web", http.Dir("./web"))
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

	fmt.Println(readme)
	fmt.Println("easyNmon running...")
	r.Run(":" + *port) // listen and serve on 0.0.0.0:8080
}

func start(c *gin.Context) { // 格式 ?n=name&t=time 其中&后可为空 默认30分钟
	name := c.DefaultQuery("n", "name")  // 取name值
	timeStr := c.DefaultQuery("t", "30") // 取执行时间,单位分钟
	fileName := name + time.Now().Format("20060102150405")

	t, _ := strconv.Atoi(timeStr)
	if t == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": string("执行错误，请检查参数是否正确！"),
		})
		return
	}
	go func() {
		fp := filepath.Join(ReportDir, fileName)
		os.MkdirAll(fp, 777)
		exec.Command("cp", "-f", "web/js/echarts.min.js", fp).Run()
		exec.Command("/bin/bash", "-c", NmonPath+" -f -t -s "+timeStr+" -c 60 -m "+fp+" -F "+name).Run()
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
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "已结束EasyNmon服务!",
	})
	go func() {
		getAllReport()
		killNmon()
		os.Exit(0)
	}()
}

func stop(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "已结束所有服务器监控任务!",
	})
	go func() {
		getAllReport()
		killNmon()
	}()
}

//重新生成所有报告
func getAllReport() {
	list := getDirList(ReportDir)
	for _, v := range list {
		internal.GetNmonReport(filepath.Join(ReportDir, v), v[:len(v)-14])
	}
}

//获取文件夹列表
func getDirList(dirpath string) []string {
	var dir_list []string
	filepath.Walk(dirpath,
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
	return dir_list
}

//杀掉所有nmon进程
func killNmon() {
	ret := exec.Command("pidof", NmonPath)
	buf, _ := ret.Output()
	pids := strings.Split(strings.Replace(string(buf), "\n", "", -1), " ")
	for _, value := range pids {
		pid, _ := strconv.Atoi(value)
		syscall.Kill(pid, syscall.SIGKILL)
	}
}
